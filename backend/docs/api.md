# building-ac-3d Backend API

Base URL: `http://localhost:8080/api`

## 通用约定

所有接口返回统一 JSON 格式：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

失败响应：

```json
{
  "code": 400,
  "message": "错误信息",
  "data": null
}
```

认证接口登录成功后返回开发版 token。调用需要登录态的接口时，在 Header 中携带：

```http
Authorization: Bearer <token>
```

前端也会按当前租户附带：

```http
X-Tenant-Id: 1
```

## 数据库健康检查

### GET `/db/health`

检查 MySQL 连接是否可用。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "status": "ok"
  }
}
```

## 认证

### POST `/auth/login`

登录。租户和用户必须从数据库查询，当前阶段密码明文对比。

请求体：

```json
{
  "tenantCode": "default",
  "username": "admin",
  "password": "123456"
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "token": "dev-token-1-1-1780365834",
    "refreshToken": "dev-refresh-1-1-1780365834"
  }
}
```

### GET `/auth/me`

获取当前登录用户信息。

请求 Header：

```http
Authorization: Bearer dev-token-1-1-1780365834
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "admin",
    "nickname": "系统管理员",
    "tenant": {
      "id": 1,
      "name": "默认租户",
      "code": "default"
    },
    "roles": ["super_admin"],
    "permissions": ["*:*:*"]
  }
}
```

### POST `/auth/logout`

退出登录。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

## 租户管理

Tenant 字段：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | number | 租户 ID |
| `name` | string | 租户名称 |
| `code` | string | 租户编码，唯一 |
| `status` | number | 状态，`1` 启用，`2` 停用 |
| `remark` | string | 备注 |
| `createdAt` | string | 创建时间 |

### GET `/system/tenants`

查询租户列表。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "name": "默认租户",
      "code": "default",
      "status": 1,
      "remark": "系统初始化租户",
      "createdAt": "2026-06-01 17:20:18"
    }
  ]
}
```

### POST `/system/tenants`

新增租户。

请求体：

```json
{
  "name": "测试租户",
  "code": "test",
  "status": 1,
  "remark": "备注"
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 2
  }
}
```

### PUT `/system/tenants`

更新租户。

请求体：

```json
{
  "id": 2,
  "name": "测试租户",
  "code": "test",
  "status": 1,
  "remark": "备注"
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

### DELETE `/system/tenants?id=2`

删除租户。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

## 用户管理

User 字段：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | number | 用户 ID |
| `tenantId` | number | 所属租户 ID |
| `username` | string | 用户名 |
| `password` | string | 密码，仅创建/更新时提交；更新为空则不修改 |
| `nickname` | string | 昵称 |
| `roleIds` | number[] | 绑定角色 ID 列表 |
| `status` | number | 状态，`1` 启用，`2` 停用 |
| `createdAt` | string | 创建时间 |

### GET `/system/users`

查询用户列表。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "tenantId": 1,
      "username": "admin",
      "nickname": "系统管理员",
      "roleIds": [1],
      "status": 1,
      "createdAt": "2026-06-01 17:20:18"
    }
  ]
}
```

### POST `/system/users`

新增用户。`password` 为空时默认使用 `123456`。

请求体：

```json
{
  "tenantId": 1,
  "username": "demo",
  "password": "123456",
  "nickname": "演示用户",
  "roleIds": [1],
  "status": 1
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 2
  }
}
```

### PUT `/system/users`

更新用户。`password` 为空字符串时不修改原密码。

请求体：

```json
{
  "id": 2,
  "tenantId": 1,
  "username": "demo",
  "password": "",
  "nickname": "演示用户",
  "roleIds": [1],
  "status": 1
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

### DELETE `/system/users?id=2`

删除用户，同时删除用户角色绑定。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

## 角色管理

Role 字段：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | number | 角色 ID |
| `tenantId` | number | 所属租户 ID |
| `name` | string | 角色名称 |
| `code` | string | 角色编码，同租户内唯一 |
| `menuIds` | number[] | 绑定菜单 ID 列表 |
| `status` | number | 状态，`1` 启用，`2` 停用 |
| `createdAt` | string | 创建时间 |

### GET `/system/roles`

查询角色列表。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "tenantId": 1,
      "name": "超级管理员",
      "code": "super_admin",
      "menuIds": [1, 2, 3, 4, 5, 6],
      "status": 1,
      "createdAt": "2026-06-01 18:11:20"
    }
  ]
}
```

### POST `/system/roles`

新增角色。

请求体：

```json
{
  "tenantId": 1,
  "name": "运营管理员",
  "code": "operation_admin",
  "menuIds": [1, 2, 3],
  "status": 1
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 2
  }
}
```

### PUT `/system/roles`

更新角色。

请求体：

```json
{
  "id": 2,
  "tenantId": 1,
  "name": "运营管理员",
  "code": "operation_admin",
  "menuIds": [1, 2, 3],
  "status": 1
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

### DELETE `/system/roles?id=2`

删除角色，同时删除角色菜单绑定和用户角色绑定。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

## 菜单管理

Menu 字段：

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| `id` | number | 菜单 ID |
| `parentId` | number | 父级菜单 ID，根节点为 `0` |
| `title` | string | 菜单名称 |
| `path` | string | 路由地址 |
| `component` | string | 前端组件路径 |
| `permission` | string | 权限标识 |
| `type` | string | 菜单类型：`catalog`、`menu`、`button` |
| `sort` | number | 排序 |
| `visible` | boolean | 是否显示 |

### GET `/system/menus`

查询菜单列表。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": [
    {
      "id": 1,
      "parentId": 0,
      "title": "系统管理",
      "path": "/system",
      "component": "Layout",
      "permission": "",
      "type": "catalog",
      "sort": 1,
      "visible": true
    }
  ]
}
```

### POST `/system/menus`

新增菜单。

请求体：

```json
{
  "parentId": 1,
  "title": "审计日志",
  "path": "/system/audit-log",
  "component": "system/audit-log/index",
  "permission": "system:audit:list",
  "type": "menu",
  "sort": 10,
  "visible": true
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 7
  }
}
```

### PUT `/system/menus`

更新菜单。

请求体：

```json
{
  "id": 7,
  "parentId": 1,
  "title": "审计日志",
  "path": "/system/audit-log",
  "component": "system/audit-log/index",
  "permission": "system:audit:list",
  "type": "menu",
  "sort": 10,
  "visible": true
}
```

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

### DELETE `/system/menus?id=7`

删除菜单，同时删除角色菜单绑定。

响应示例：

```json
{
  "code": 0,
  "message": "success",
  "data": true
}
```

## 默认初始化数据

启动时通过 GORM AutoMigrate 创建以下表：

- `tenants`
- `users`
- `roles`
- `menus`
- `user_roles`
- `role_menus`
- `visual_configs`

默认账号：

| 项 | 值 |
| --- | --- |
| 租户编码 | `default` |
| 用户名 | `admin` |
| 密码 | `123456` |
| 角色 | `super_admin` |
| 权限 | `*:*:*` |
