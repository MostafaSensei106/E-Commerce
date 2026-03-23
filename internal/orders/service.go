package orders

import (
	"context"
	"log"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
)

type Service interface {
	GetOrderWithItems(ctx context.Context, id int64) ([]repo.GetOrderWithItemsRow, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) GetOrderWithItems(ctx context.Context, id int64) ([]repo.GetOrderWithItemsRow, error) {
	orderItems, err := s.repo.GetOrderWithItems(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return orderItems, nil
}
