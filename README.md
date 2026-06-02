# building-ac-3d

楼宇空调 3D 可视化配置项目。

## 目录结构

```txt
backend/   Go 后端服务
frontend/  Vue + Vite 前端页面
```

## 1. 创建 MySQL 数据库

你的 MySQL 连接信息：

```txt
主机：localhost
端口：3306
用户名：root
数据库：building_ac_3d
```

先在本机 MySQL 中创建数据库：

```sql
CREATE DATABASE building_ac_3d
  DEFAULT CHARACTER SET utf8mb4
  DEFAULT COLLATE utf8mb4_unicode_ci;
```

后端启动时会自动创建 `visual_configs` 表。

## 2. 配置后端环境变量

进入后端目录：

```bash
cd backend
cp .env.example .env
```

然后编辑 `backend/.env`：

```env
APP_PORT=8080

MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DATABASE=building_ac_3d
MYSQL_USER=root
MYSQL_PASSWORD=你的 MySQL 密码

CORS_ALLOW_ORIGIN=http://localhost:5173
```

注意：`.env` 是本地私密配置，已经被 `.gitignore` 忽略，不要提交到 GitHub。

## 3. 启动后端

```bash
cd backend
go mod tidy
go run .
```

后端默认运行在：

```txt
http://localhost:8080
```

可以访问下面接口测试 MySQL 是否连接成功：

```txt
GET http://localhost:8080/api/db/health
```

## 4. 启动前端

打开另一个终端：

```bash
cd frontend
npm install
npm run dev
```

前端默认运行在：

```txt
http://localhost:5173
```

前端已经配置 Vite 代理，页面请求 `/api` 会自动转发到 `http://localhost:8080`。

## 5. 当前已提供接口

### 检测数据库连接

```txt
GET /api/db/health
```

### 查询 3D 可视化配置列表

```txt
GET /api/visual-configs
```

### 新增 3D 可视化配置

```txt
POST /api/visual-configs
Content-Type: application/json

{
  "name": "测试配置",
  "description": "用于保存楼宇、楼层、设备、点位绑定关系",
  "configData": {
    "building": {
      "name": "示例楼宇"
    },
    "devices": []
  }
}
```
