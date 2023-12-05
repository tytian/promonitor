
CREATE DATABASE IF NOT EXISTS monitor
       DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

use monitor;
SET NAME utf8;

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `age` int(3) NOT NULL,
    `birthday` datetime NOT NULL,
    `email` varchar(50) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime ,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;