package orders

import repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"

type Service interface {
	GetAllOrders() ([]repo.Order, error)
	GetOrderByID(id int64) (repo.Order, error)
}
