CREATE TABLE IF NOT EXISTS orders (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    product_id INT,
    quantity INT,
    shipping_address VARCHAR(255),
    request_id UUID UNIQUE
)