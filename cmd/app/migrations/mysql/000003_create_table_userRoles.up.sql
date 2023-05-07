CREATE TABLE `userRoles`
(
    `id`      varchar(255) NOT NULL,
    `userId`  varchar(255) NOT NULL,
    `roleId`  varchar(255) NOT NULL,
    `created` timestamp    NOT NULL,
    `updated` timestamp,
    `deleted` timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4;