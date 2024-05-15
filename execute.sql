USE mydatabase;

CREATE TABLE products (
    uuid VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    PRIMARY KEY (uuid)
);