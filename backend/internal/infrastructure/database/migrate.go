package database

import "database/sql"

func Migrate(db *sql.DB) error {
    statements := []string{
        `CREATE TABLE IF NOT EXISTS tenants (
            id BIGINT PRIMARY KEY AUTO_INCREMENT,
            name VARCHAR(100) NOT NULL,
            code VARCHAR(64) NOT NULL UNIQUE,
            status TINYINT NOT NULL DEFAULT 1,
            remark VARCHAR(255) NOT NULL DEFAULT '',
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
        `CREATE TABLE IF NOT EXISTS users (
            id BIGINT PRIMARY KEY AUTO_INCREMENT,
            tenant_id BIGINT NOT NULL,
            username VARCHAR(64) NOT NULL,
            password VARCHAR(255) NOT NULL,
            nickname VARCHAR(64) NOT NULL DEFAULT '',
            status TINYINT NOT NULL DEFAULT 1,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            UNIQUE KEY uk_tenant_username (tenant_id, username)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
        `CREATE TABLE IF NOT EXISTS roles (
            id BIGINT PRIMARY KEY AUTO_INCREMENT,
            tenant_id BIGINT NOT NULL,
            name VARCHAR(64) NOT NULL,
            code VARCHAR(64) NOT NULL,
            status TINYINT NOT NULL DEFAULT 1,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
        `CREATE TABLE IF NOT EXISTS menus (
            id BIGINT PRIMARY KEY AUTO_INCREMENT,
            parent_id BIGINT NOT NULL DEFAULT 0,
            title VARCHAR(64) NOT NULL,
            path VARCHAR(255) NOT NULL DEFAULT '',
            component VARCHAR(255) NOT NULL DEFAULT '',
            permission VARCHAR(128) NOT NULL DEFAULT '',
            type VARCHAR(16) NOT NULL,
            sort INT NOT NULL DEFAULT 0,
            visible TINYINT NOT NULL DEFAULT 1
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
        `CREATE TABLE IF NOT EXISTS visual_configs (
            id BIGINT PRIMARY KEY AUTO_INCREMENT,
            tenant_id BIGINT NOT NULL DEFAULT 1,
            name VARCHAR(100) NOT NULL,
            description VARCHAR(255) NOT NULL DEFAULT '',
            config_data JSON NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
        `INSERT IGNORE INTO tenants (id, name, code, status, remark) VALUES (1, '默认租户', 'default', 1, '系统初始化租户');`,
        `INSERT IGNORE INTO users (id, tenant_id, username, password, nickname, status) VALUES (1, 1, 'admin', '123456', '系统管理员', 1);`,
    }

    for _, statement := range statements {
        if _, err := db.Exec(statement); err != nil {
            return err
        }
    }

    return nil
}
