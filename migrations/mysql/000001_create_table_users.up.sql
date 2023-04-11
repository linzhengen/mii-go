CREATE TABLE `users`
(
    `id`       varchar(255) NOT NULL,
    `name`     varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `email`    varchar(255) NOT NULL,
    `status`   varchar(255) NOT NULL,
    `created`  timestamp    NOT NULL,
    `updated`  timestamp NULL,
    `deleted`  timestamp NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;