CREATE TABLE `roles`
(
    `id`       varchar(255)  NOT NULL,
    `name`     varchar(255)  NOT NULL,
    `apiGroup` varchar(255)  NOT NULL, -- endpoint, menu,
    `resource` varchar(1024) NOT NULL, -- POST./rest/v1/users, GET./rest/v1/users/{id}, User, Role, Menu
    `created`  timestamp     NOT NULL,
    `updated`  timestamp,
    `deleted`  timestamp,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARACTER SET utf8mb4;