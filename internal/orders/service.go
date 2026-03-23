package orders

import (
	"context"
	"errors"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrProductHasNoStock = errors.New("product has no enought stock")
)

type Service interface {
	PlaceNewOrder(ctx context.Context, params repo.CreateOrderParams) (repo.Order, error)
}

type svc struct {
	repo repo.Querier
	db   *pgx.Conn
}

func NewService(repo repo.Querier, db *pgx.Conn) Service {
	return &svc{
		repo: repo,
		db:   db,
	}
}

func (s *svc) PlaceNewOrder(ctx context.Context, p repo.CreateOrderParams) (repo.Order, error) {

	order, err := s.repo.CreateOrder(ctx, p)
	if err != nil {
		return repo.Order{}, err
	}

	for _, item := range p.Items {
		product, err := s.repo.GetProductByID(ctx, item.ProductID)
		if err != nil {
			return repo.Order{}, ErrProductNotFound
		}

		if product.Quantity < item.Quantity {
			return repo.Order{}, ErrProductHasNoStock
		}

		_, err = s.repo.CreateOrderItem(ctx, repo.CreateOrderItemParams{
			OrderID:      order.ID,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PriceInCents: product.PriceInCents,
		})
		if err != nil {
			return repo.Order{}, err
		}
	}
	return order, nil
}
