package client

import (
	context "context"
	"fmt"
	"io/ioutil"
	"kubectl-cli/util"
	"os"
	"time"

	common "kubectl-cli/common"

	b64 "encoding/base64"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type ClientServer struct {
	UnimplementedKubectlClientServer
}

func makeKubeConfigFile(kubeConfigContent string) (*rest.Config, error) {
	sDec, _ := b64.StdEncoding.DecodeString(kubeConfigContent)
	fileName := util.RandStringRunes(32)
	ioutil.WriteFile(fileName, sDec, 0644)
	config, err := clientcmd.BuildConfigFromFlags("", fileName)
	defer os.Remove(fileName)
	if err != nil {
		return nil, err
	} else {
		return config, nil
	}
}

func getClient(kubeConfigContent string) (*kubernetes.Clientset, error) {
	config, err := makeKubeConfigFile(kubeConfigContent)

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func (s *ClientServer) GetServices(ctx context.Context, in *GetServiceRequest) (*GetServiceResponse, error) {
	if in.Req.Kubeconfig == "" {
		return &GetServiceResponse{Resp: &common.CommonResponse{Descryption: "empty kubeconfig", ResultCode: 1}}, nil
	}

	clientset, err := getClient(in.Req.Kubeconfig)

	if err != nil {
		return nil, err
	}

	services, err := clientset.CoreV1().Services(in.Namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	ServiceInfoArr := make([]*ServiceInfo, 0)
	for _, service := range services.Items {
		fmt.Printf("%+v\n", service.Spec)
		se := labels.Set(service.Spec.Selector)
		fmt.Printf("%s\n", se.AsSelector().String())

		serviceCreationTime := service.GetCreationTimestamp()
		age := time.Since(serviceCreationTime.Time).Round(time.Second)
		serviceInfo := ServiceInfo{Name: service.Name, Type: string(service.Spec.Type), ExternalIp: service.Spec.ExternalIPs, Age: age.String(), ClusterIp: service.Spec.ClusterIP, LabelSelector: se.AsSelector().String()}

		portInfoArr := make([]*PortInfo, 0)
		for _, portInfo := range service.Spec.Ports {
			portInfo := PortInfo{Port: portInfo.Port, TargetPort: portInfo.TargetPort.IntVal, Protocol: string(portInfo.Protocol)}
			portInfoArr = append(portInfoArr, &portInfo)
		}
		serviceInfo.Ports = portInfoArr
		ServiceInfoArr = append(ServiceInfoArr, &serviceInfo)
	}
	if err != nil {
		return nil, err
	}

	resp := &GetServiceResponse{Services: ServiceInfoArr, Resp: &common.CommonResponse{Descryption: "get services successful", ResultCode: 0}}
	return resp, nil
}

func (s *ClientServer) GetNamespaces(ctx context.Context, in *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	if in.Req.Kubeconfig == "" {
		return &GetNamespaceResponse{Resp: &common.CommonResponse{Descryption: "empty kubeconfig", ResultCode: 1}}, nil
	}

	clientset, err := getClient(in.Req.Kubeconfig)

	if err != nil {
		return nil, err
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	namespaceArr := make([]string, 0)

	for _, namespace := range namespaces.Items {
		namespaceArr = append(namespaceArr, namespace.Name)
	}

	resp := &GetNamespaceResponse{Namespace: namespaceArr, Resp: &common.CommonResponse{Descryption: "get namespaces successful", ResultCode: 0}}
	return resp, nil
}

func (s *ClientServer) GetPods(ctx context.Context, in *GetPodsRequest) (*GetPodsResponse, error) {
	if in.Req.Kubeconfig == "" {
		return &GetPodsResponse{Resp: &common.CommonResponse{Descryption: "empty kubeconfig", ResultCode: 1}}, nil
	}

	clientset, err := getClient(in.Req.Kubeconfig)

	if err != nil {
		return nil, err
	}

	namespace := in.Namespace

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod in namespace %s: %v\n",
			namespace, statusError.ErrStatus.Message)
		return nil, statusError
	} else if err != nil {
		return nil, err
	}

	podInfoArr := make([]*PodInfo, 0)
	for _, pod := range pods.Items {
		// Calculate the age of the pod
		podInfo := PodInfo{}
		podCreationTime := pod.GetCreationTimestamp()
		age := time.Since(podCreationTime.Time).Round(time.Second)

		podStatus := pod.Status
		if len(pod.Spec.Containers) == 1 && podStatus.ContainerStatuses[0].State.Waiting != nil {
			fmt.Printf("%+v\n", podStatus.ContainerStatuses[0])
			podInfo.Status = podStatus.ContainerStatuses[0].State.Waiting.Reason
		} else {
			podInfo.Status = fmt.Sprintf("%v", podStatus.Phase)
		}

		var containerRestarts int32
		var containerReady int
		var totalContainers int

		// If a pod has multiple containers, get the status from all
		for container := range pod.Spec.Containers {
			containerRestarts += podStatus.ContainerStatuses[container].RestartCount
			if podStatus.ContainerStatuses[container].Ready {
				containerReady++
			}
			totalContainers++
		}
		podInfo.Age = age.String()
		podInfo.Restarts = containerRestarts
		podInfo.Name = pod.GetName()
		podInfo.ReadyPods = int32(containerReady)
		podInfo.TotalPods = int32(totalContainers)
		podInfoArr = append(podInfoArr, &podInfo)
	}

	fmt.Printf("%+v\n", podInfoArr)
	return &GetPodsResponse{Resp: &common.CommonResponse{Descryption: "get pods successful", ResultCode: 0}, Info: podInfoArr}, nil
}
