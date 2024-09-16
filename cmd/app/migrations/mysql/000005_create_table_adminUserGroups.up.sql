CREATE TABLE adminUserGroups
(
    id          CHAR(36)     NOT NULL PRIMARY KEY, -- UUID
    groupName   VARCHAR(255) NOT NULL,
    status      ENUM('Active', 'Inactive') NOT NULL,
    description TEXT,
    updatedAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    createdAt   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;