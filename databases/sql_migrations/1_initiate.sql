-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by INT
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by INT
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    description TEXT,
    image_url TEXT,
    release_year INT,
    price INT,
    total_page INT,
    thickness VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by INT,
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modified_by INT,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);


-- +migrate StatementEnd