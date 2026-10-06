CREATE DATABASE IF NOT EXISTS secure_login_db;

USE secure_login_db;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role ENUM('user', 'admin') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (email, password_hash, role)
VALUES
(
    'user@gmail.com',
    '$2a$10$RPgP7viLD/HorGRHxrDOFVhJF8SUP9G/sqZNmfCj0u/.',
    'user'
),
(
    'admin@gmail.com',
    '$2a$10$eNbl37oL8wU9ytpH4j9ORZzt.Vk/bJ36w13JwJPd7tk1.',
    'admin'
);