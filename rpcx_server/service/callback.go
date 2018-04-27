package service

import (
	"context"
	"fmt"
)

type CallbackRequest struct {
	info string
}

type CallbackResponse struct {
	info string
}

func (s *Service) Callback(ctx context.Context, request *CallbackRequest, reply *CallbackResponse) error {
	fmt.Println("service Callback")
	return nil
}
