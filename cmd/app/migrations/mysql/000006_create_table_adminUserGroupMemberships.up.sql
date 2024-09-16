CREATE TABLE adminUserGroupMemberships
(
    id               CHAR(36) NOT NULL PRIMARY KEY, -- UUID
    adminUserId      CHAR(36) NOT NULL,             -- Foreign Key to adminUsers.id
    adminUserGroupId CHAR(36) NOT NULL,             -- Foreign Key to adminUserGroups.id
    updatedAt        TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    createdAt        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;