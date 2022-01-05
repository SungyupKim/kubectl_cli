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

func (s *ClientServer) GetNamespaces(ctx context.Context, in *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	if in.Req.Kubeconfig == "" {
		return &GetNamespaceResponse{Resp: &common.CommonResponse{Descryption: "empty kubeconfig", ResultCode: 1}}, nil
	}

	//config, err := makeKubeConfigFile(in.Req.Kubeconfig)

	//	if err != nil {
	//		return nil, err
	//	}

	return nil, nil
}

func (s *ClientServer) GetPods(ctx context.Context, in *GetPodsRequest) (*GetPodsResponse, error) {
	if in.Req.Kubeconfig == "" {
		return &GetPodsResponse{Resp: &common.CommonResponse{Descryption: "empty kubeconfig", ResultCode: 1}}, nil
	}

	config, err := makeKubeConfigFile(in.Req.Kubeconfig)

	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	namespace := in.Namespace
	pod := "example-xxxxx"
	pods, err = clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
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
