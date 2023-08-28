CREATE TABLE product_variants (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    updater_id UUID,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    price INTEGER NOT NULL,
    product_id UUID NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    variant JSONB NOT NULL,
    sku VARCHAR(255) NOT NULL,
    stock INTEGER NOT NULL DEFAULT 0
);