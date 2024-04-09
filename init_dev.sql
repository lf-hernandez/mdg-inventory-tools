-- Ensure the UUID extension is available
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Ensure the pgcrypto extension is available
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create the 'inventory' table with metadata columns
CREATE TABLE inventory (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'item' table with metadata columns
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
    inventory_id UUID REFERENCES inventory(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for 'item' table
CREATE INDEX idx_item_part_number ON item (part_number);
CREATE INDEX idx_item_purchase_order ON item (purchase_order);

-- Create the 'app_user' table with metadata columns
CREATE TABLE app_user (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Seed user
INSERT INTO app_user (name, email, password)
VALUES ('Test User', 'test1@app.com', crypt('123', gen_salt('bf')));

-- Seed items
INSERT INTO item (part_number, serial_number, purchase_order, description, category, price, quantity, status, repair_order_number, condition)
VALUES
('PN-001', 'SN-001', 'PO-001', 'Item Description 1', 'Category 1', 100.00, 5, 'Available', 'RO-001', 'New'),
('PN-002', 'SN-002', 'PO-002', 'Item Description 2', 'Category 2', 200.00, 3, 'In Use', 'RO-002', 'Used'),
('PN-003', 'SN-003', 'PO-003', 'Item Description 3', 'Category 3', 150.00, 2, 'Maintenance', 'RO-003', 'Refurbished');

-- Create a function for updating the 'modified_at' column
CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create triggers for each table to automatically update 'modified_at'
CREATE TRIGGER update_inventory_modtime
BEFORE UPDATE ON inventory
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_item_modtime
BEFORE UPDATE ON item
FOR EACH ROW EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_app_user_modtime
BEFORE UPDATE ON app_user
FOR EACH ROW EXECUTE FUNCTION update_modified_column();
