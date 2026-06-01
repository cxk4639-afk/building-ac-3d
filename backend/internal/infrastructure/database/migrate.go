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
            updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            UNIQUE KEY uk_tenant_role_code (tenant_id, code)
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
        `CREATE TABLE IF NOT EXISTS user_roles (
            user_id BIGINT NOT NULL,
            role_id BIGINT NOT NULL,
            PRIMARY KEY (user_id, role_id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
        `CREATE TABLE IF NOT EXISTS role_menus (
            role_id BIGINT NOT NULL,
            menu_id BIGINT NOT NULL,
            PRIMARY KEY (role_id, menu_id)
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
        `INSERT IGNORE INTO roles (id, tenant_id, name, code, status) VALUES (1, 1, '超级管理员', 'super_admin', 1);`,
        `INSERT IGNORE INTO menus (id, parent_id, title, path, component, permission, type, sort, visible) VALUES
            (1, 0, '系统管理', '/system', 'Layout', '', 'catalog', 1, 1),
            (2, 1, '租户管理', '/system/tenant', 'system/tenant/index', 'system:tenant:list', 'menu', 1, 1),
            (3, 1, '用户管理', '/system/user', 'system/user/index', 'system:user:list', 'menu', 2, 1),
            (4, 1, '角色管理', '/system/role', 'system/role/index', 'system:role:list', 'menu', 3, 1),
            (5, 1, '菜单管理', '/system/menu', 'system/menu/index', 'system:menu:list', 'menu', 4, 1),
            (6, 0, '3D配置', '/visual-config', 'visual-config/index', 'visual:config:list', 'menu', 2, 1);`,
        `INSERT IGNORE INTO user_roles (user_id, role_id) VALUES (1, 1);`,
        `INSERT IGNORE INTO role_menus (role_id, menu_id) VALUES (1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6);`,
    }

    for _, statement := range statements {
        if _, err := db.Exec(statement); err != nil {
            return err
        }
    }

    return nil
}
