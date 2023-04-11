CREATE TABLE `roles`
(
    `id`      varchar(255) NOT NULL,
    `name`    varchar(255) NOT NULL,
    `status`  varchar(255) NOT NULL,
    `created` timestamp    NOT NULL,
    `updated` timestamp,
    `deleted` timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;