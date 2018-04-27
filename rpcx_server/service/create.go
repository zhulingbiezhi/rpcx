package service

import (
	"context"
	"fmt"
)

type CreateRequest struct {
	Info string `json:"info"`
}
type CreateResponse struct {
	Info string `json:"info"`
}

func (s *Service) Create(ctx context.Context, request *CreateRequest, reply *CreateResponse) error {
	fmt.Println("service create")
	reply.Info = "hello"
	return nil
}
