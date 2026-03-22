package products

import "context"

type Service interface {
	GatAll(ctx context.Context) error
}

type svc struct {
	// repository
}

func NewService() Service {
	return &svc{}
}

func (s *svc) GatAll(ctx context.Context) error {
	return nil
}
