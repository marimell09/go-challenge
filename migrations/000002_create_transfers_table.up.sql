CREATE TABLE IF NOT EXISTS transfers(
    id SERIAL PRIMARY KEY,
    account_origin_id INT NOT NULL,
    account_destination_id INT NOT NULL,
    amount NUMERIC(10, 2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);