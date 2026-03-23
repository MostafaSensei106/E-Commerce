package orders

import (
	"context"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
)

type Service interface {
	PlaceNewOrder(ctx context.Context, params repo.CreateOrderParams) (repo.Order, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) PlaceNewOrder(ctx context.Context, p repo.CreateOrderParams) (repo.Order, error) {

}
