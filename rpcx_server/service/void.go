package service

import (
	"context"
	"fmt"
)

type VoidRequest struct {
}
type VoidResponse struct {
}

func (s *Service) Void(ctx context.Context, request *VoidRequest, reply *VoidResponse) error {
	fmt.Println("service Void")
	return nil
}
