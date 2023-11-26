CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL
);

CREATE TABLE item (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_number TEXT NOT NULL,
    serial_number TEXT,
    purchase_order TEXT,
    description TEXT,
    category TEXT,
    price NUMERIC(10, 2),
    quantity INTEGER,
    inventory_id UUID REFERENCES inventory(inventory_id)
);

CREATE INDEX idx_item_part_number ON item (part_number);
CREATE INDEX idx_item_purchase_order ON item (purchase_order);

CREATE TABLE app_user (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);
