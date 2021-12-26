package client

import context "context"

type ClientServer struct {
	UnimplementedKubectlClientServer
}

func (s *ClientServer) Get(ctx context.Context, in *GetPodsRequest) (*GetPodsResponse, error) {
	return &GetPodsResponse{}, nil
}
