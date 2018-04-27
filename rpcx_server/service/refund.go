package service

import (
	"context"
	"fmt"
)

type RefundRequest struct {
}
type RefundResponse struct {
}

func (s *Service) Refund(ctx context.Context, request *RefundRequest, reply *RefundResponse) error {
	fmt.Println("service create")
	return nil
}
