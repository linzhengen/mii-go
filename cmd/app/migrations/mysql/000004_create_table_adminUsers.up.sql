CREATE TABLE adminUsers
(
    id           CHAR(36) PRIMARY KEY,                                            -- UUID stored as a CHAR(36)
    userName     VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL UNIQUE,                                    -- Unique constraint on email
    passwordHash VARCHAR(255) NOT NULL,
    status       ENUM('Active', 'Inactive') DEFAULT 'Active',                     -- Status with two possible values
    updatedAt    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Auto-updates on row update
    createdAt    TIMESTAMP DEFAULT CURRENT_TIMESTAMP                              -- Auto-sets on row creation
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;