package products

import (
	"context"
	"errors"
	"log"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type Service interface {
	GetAllProducts(ctx context.Context) ([]repo.Product, error)
	GetProductByID(ctx context.Context, id int64) (repo.Product, error)
	CreateNewProduct(ctx context.Context, params repo.CreateProductParams) (repo.Product, error)
	UpdateProduct(ctx context.Context, params repo.UpdateProductWhereIDParams) (repo.Product, error)
	DeleteProduct(ctx context.Context, id int64) error
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

func (s *svc) CreateNewProduct(ctx context.Context, p repo.CreateProductParams) (repo.Product, error) {
	product, err := s.repo.CreateProduct(ctx, p)
	if err != nil {
		log.Println(err.Error())
		return repo.Product{}, err
	}
	return product, nil
}

func (s *svc) UpdateProduct(ctx context.Context, p repo.UpdateProductWhereIDParams) (repo.Product, error) {
	product, err := s.repo.UpdateProductWhereID(ctx, p)
	if err != nil {
		log.Println(err.Error())
		return product, err
	}
	return product, nil
}

func (s *svc) DeleteProduct(ctx context.Context, id int64) error {
	exists, err := s.repo.ProductExists(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if exists == false {
		return ErrProductNotFound
	}

	if exists == true {
		err := s.repo.DeleteProduct(ctx, id)
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}
	return nil
}
