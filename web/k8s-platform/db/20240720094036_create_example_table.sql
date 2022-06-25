-- +goose Up
CREATE TABLE `example` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `name` varchar(32) COLLATE utf8mb4_general_ci NOT NULL,
                            `num` int NOT NULL,
                            `created_at` datetime DEFAULT NULL,
                            `updated_at` datetime DEFAULT NULL,
                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- +goose Down
DROP TABLE example;
