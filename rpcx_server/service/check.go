package service

import (
	"context"
	"fmt"
)

type CheckRequest struct {
}

type CheckResponse struct {
}

func (s *Service) Check(ctx context.Context, request *CheckRequest, reply *CheckResponse) error {
	fmt.Println("service Check")
	return nil
}
