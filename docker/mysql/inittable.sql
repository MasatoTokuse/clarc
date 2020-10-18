CREATE TABLE IF NOT EXISTS users (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `nickname` varchar(255) NOT NULL,
    PRIMARY KEY(id)
);