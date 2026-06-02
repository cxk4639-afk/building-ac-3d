# Go 后端热编译工具 Air 使用笔记

Air 是 Go 项目开发时常用的热编译 / 热重启工具。

它的作用是：

```txt
修改 Go 代码
  ↓
Air 自动监听文件变化
  ↓
自动重新编译
  ↓
自动重启后端服务
```

这样开发后端时就不用每次手动执行：

```bash
go run .
```

## 1. 安装 Air

安装命令在哪个目录执行都可以，因为 Air 会安装到 Go 的全局工具目录。

推荐在任意目录执行：

```bash
go install github.com/air-verse/air@latest
```

安装完成后检查版本：

```bash
air -v
```

如果提示：

```txt
zsh: command not found: air
```

说明 Go 的工具目录还没有加入 PATH。

可以先直接这样验证 Air 是否已经安装成功：

```bash
$(go env GOPATH)/bin/air -v
```

如果能看到版本号，就说明安装成功。

## 2. 配置 PATH

把 Go 工具目录加入 zsh 配置：

```bash
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
source ~/.zshrc
```

然后重新检查：

```bash
air -v
```

## 3. 初始化 Air 配置

运行 Air 时，需要进入后端目录。

本项目后端目录是：

```bash
cd /Users/chenxk/GolandProjects/awesomeProject/backend
```

第一次使用 Air，可以初始化配置文件：

```bash
air init
```

执行后会生成：

```txt
.air.toml
```

## 4. 推荐的 .air.toml 配置

如果需要手动配置，可以使用下面内容：

```toml
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/building-ac-3d-backend ."
bin = "./tmp/building-ac-3d-backend"
include_ext = ["go", "tpl", "tmpl", "html"]
exclude_dir = ["tmp", "vendor", "frontend", "node_modules"]
delay = 1000
stop_on_error = true

[log]
time = true

[misc]
clean_on_exit = true
```

## 5. 启动后端热编译

以后启动后端，不用再执行：

```bash
go run .
```

改成：

```bash
air
```

如果 PATH 还没有生效，也可以临时用完整路径启动：

```bash
$(go env GOPATH)/bin/air
```

启动成功后会看到类似：

```txt
building-ac-3d backend is running at http://localhost:8080
```

## 6. 日常开发启动方式

前端：

```bash
cd /Users/chenxk/GolandProjects/awesomeProject/frontend
npm run dev
```

后端：

```bash
cd /Users/chenxk/GolandProjects/awesomeProject/backend
air
```

## 7. 常见问题

### air: command not found

原因：Air 已安装，但 Go 工具目录没有加入 PATH。

解决：

```bash
echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
source ~/.zshrc
```

或者临时执行：

```bash
$(go env GOPATH)/bin/air
```

### 修改代码后没有自动重启

检查当前目录是否是后端目录：

```bash
pwd
```

必须在：

```txt
/Users/chenxk/GolandProjects/awesomeProject/backend
```

或者你的实际项目 backend 目录。

### 后端启动失败

先确认 `.env` 是否存在：

```bash
ls -la
```

如果没有：

```bash
cp .env.example .env
```

并确认 MySQL 配置正确：

```env
APP_PORT=8080
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DATABASE=building_ac_3d
MYSQL_USER=root
MYSQL_PASSWORD=你的 MySQL 密码
CORS_ALLOW_ORIGIN=http://localhost:5173
JWT_SECRET=building-ac-3d-dev-secret
```

## 8. 简短版命令

以后你只需要记住：

```bash
cd /Users/chenxk/GolandProjects/awesomeProject/backend
air
```

如果 `air` 找不到：

```bash
$(go env GOPATH)/bin/air
```
