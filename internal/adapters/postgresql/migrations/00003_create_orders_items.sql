-- +goose Up
CREATE TABLE IF NOT EXISTS orders_items (
    id BIGSERIAL PRIMARY KEY,

    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,

    quantity INTEGER NOT NULL CHECK (quantity > 0),
    price_in_cents DOUBLE PRECISION NOT NULL CHECK (price_in_cents >= 0),

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,

    UNIQUE (order_id, product_id) 
);
-- +goose Down
DROP TABLE IF EXISTS orders_items;
