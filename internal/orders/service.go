package orders

import (
	"context"
	"errors"
	"log"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5"
)

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrProductHasNoStock = errors.New("product has no enought stock")
)

type Service interface {
	PlaceNewOrder(ctx context.Context, params repo.CreateOrderParams) (repo.Order, error)
	GetAllOrders(ctx context.Context) ([]repo.Order, error)
	GetOrderByID(ctx context.Context, id int64) (repo.Order, error)
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

func (s *svc) GetAllOrders(ctx context.Context) ([]repo.Order, error) {
	orders, err := s.repo.GetAllOrders(ctx)
	if err != nil {
		log.Println(err.Error())
		return []repo.Order{}, err
	}
	return orders, nil
}

func (s *svc) GetOrderByID(ctx context.Context, id int64) (repo.Order, error) {
	order, err := s.repo.GetOrderByID(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return repo.Order{}, err
	}
	return order, nil
}

func (s *svc) PlaceNewOrder(ctx context.Context, p repo.CreateOrderParams) (repo.Order, error) {

	order, err := s.repo.CreateOrder(ctx, p)
	if err != nil {
		return repo.Order{}, err
	}

	// for _, item := range p.Items {
	// 	product, err := s.repo.GetProductByID(ctx, item.ProductID)
	// 	if err != nil {
	// 		return repo.Order{}, ErrProductNotFound
	// 	}

	// 	if product.Quantity < item.Quantity {
	// 		return repo.Order{}, ErrProductHasNoStock
	// 	}

	// 	_, err = s.repo.CreateOrderItem(ctx, repo.CreateOrderItemParams{
	// 		OrderID:      order.ID,
	// 		ProductID:    item.ProductID,
	// 		Quantity:     item.Quantity,
	// 		PriceInCents: product.PriceInCents,
	// 	})
	// 	if err != nil {
	// 		return repo.Order{}, err
	// 	}
	// 	err = s.repo.IncreaseProductQuantity(ctx, repo.IncreaseProductQuantityParams{
	// 		Quantity: -item.Quantity,
	// 		ID:       item.ProductID,
	// 	})
	// 	if err != nil {
	// 		return repo.Order{}, err
	// 	}
	// }
	return order, nil
}
