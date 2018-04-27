package service

import (
	"context"
	"fmt"
)

type CaptureRequest struct {
}

type CaptureResponse struct {
}

func (s *Service) Capture(ctx context.Context, request *CaptureRequest, reply *CaptureResponse) error {
	fmt.Println("service Capture")
	return nil
}
