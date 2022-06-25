-- +goose Up
-- SQL 在此块中运行时，将应用迁移。

CREATE TABLE users_golang (
   id INT AUTO_INCREMENT PRIMARY KEY,
   username VARCHAR(255) NOT NULL,
   password VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL 在此块中运行时，将撤销迁移。

DROP TABLE users_golang;