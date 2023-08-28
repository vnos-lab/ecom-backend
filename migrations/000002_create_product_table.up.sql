CREATE TABLE IF NOT EXISTS products (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    updater_id UUID,
    description VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    thumbnail VARCHAR(255) NOT NULL,
    images JSONB NOT NULL,
    rating INT NOT NULL DEFAULT 0,
    is_new BOOL NOT NULL DEFAULT false,
    category_id uuid,
    campaign_id uuid,
    total_sold INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
