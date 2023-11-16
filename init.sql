CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE item (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    item_id TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10, 2),
    quantity INTEGER
);

CREATE INDEX idx_item_item_id ON item (item_id);
