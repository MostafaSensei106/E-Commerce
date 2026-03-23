package cmd

import (
	"log/slog"
	"net/http"
	"time"

	repo "github.com/MostafaSensei106/E-Commerce/internal/adapters/postgresql/sqlc"
	"github.com/MostafaSensei106/E-Commerce/internal/orders"
	"github.com/MostafaSensei106/E-Commerce/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	port         string
	db           databaseConfig
	writeTimeout time.Duration
	readTimeout  time.Duration
	idleTimeout  time.Duration
}

type databaseConfig struct {
	dsn string
}

// Mount
func (a *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All Good!"))
	})

	productService := products.NewService(repo.New(a.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.GetAllProductsHandler)
	r.Post("/products", productHandler.CreateProductHandler)
	r.Get("/products/{id}", productHandler.GetProductByIDHandler)
	r.Put("/products/{id}", productHandler.UpdateProductHandler)
	r.Delete("/products/{id}", productHandler.DeleteProductHandler)

	ordersService := orders.NewService(repo.New(a.db))
	ordersHandler := orders.NewHandler(ordersService)

	r.Post("/orders", ordersHandler.PlaceNewOrderHandler)

	return r
}

// Run
func (a *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         a.config.port,
		Handler:      h,
		WriteTimeout: a.config.writeTimeout,
		ReadTimeout:  a.config.readTimeout,
		IdleTimeout:  a.config.idleTimeout,
	}

	slog.Info("Server started", "addr", a.config.port)

	return srv.ListenAndServe()
}
