CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE item (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_number TEXT NOT NULL UNIQUE,
    serial_number TEXT,
    purchase_order TEXT,
    description TEXT,
    category TEXT,
    price NUMERIC(10, 2),
    quantity INTEGER,
    status TEXT,
    repair_order_number TEXT,
    condition TEXT,
    location TEXT,
    notes TEXT,
    inventory_id UUID REFERENCES inventory(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_item_part_number ON item (part_number);
CREATE INDEX idx_item_purchase_order ON item (purchase_order);

CREATE TYPE user_role AS ENUM ('Admin', 'Manager', 'Employee');

CREATE TABLE app_user (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL,
    permissions JSONB,
    inventories TEXT[] ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_inventory_mtime
BEFORE UPDATE ON inventory
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_item_mtime
BEFORE UPDATE ON item
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_app_user_mtime
BEFORE UPDATE ON app_user
FOR EACH ROW EXECUTE FUNCTION update_modified_column();
