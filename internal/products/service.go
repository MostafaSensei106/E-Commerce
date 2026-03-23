package products

import (
	"context"
	"log"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
)

type Service interface {
	GetAllProducts(ctx context.Context) ([]repo.Product, error)
	GetProductByID(ctx context.Context, id int64) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) GetAllProducts(ctx context.Context) ([]repo.Product, error) {
	products, err := s.repo.GetAllProducts(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return products, nil
}

func (s *svc) GetProductByID(ctx context.Context, id int64) (repo.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return repo.Product{}, err
	}
	return product, nil
}

func (s *svc) CreateProduct(ctx context.Context, product repo.CreateProductParams) (repo.Product, error) {
	newProduct, err := s.repo.CreateProduct(ctx, product)
	if err != nil {
		log.Println(err.Error())
		return repo.Product{}, err
	}
	return newProduct, nil
}
