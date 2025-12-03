# donkey-ucenter-backend

## 项目说明

donkey-ucenter-backend 是一个基于 Go 语言开发的用户中心后端服务系统。项目采用 Gin 框架构建，提供用户管理、身份验证、邮箱验证等核心功能。

### 技术栈

- **语言**: Go 1.21+
- **Web 框架**: Gin
- **数据库**: MySQL (GORM)
- **缓存**: Redis
- **配置管理**: Viper (TOML)
- **JWT 认证**: golang-jwt/jwt

### 主要功能

- 用户注册、登录、认证
- 用户信息管理
- 邮箱验证
- JWT Token 认证
- 管理员功能

## 安装方式

项目提供两种安装方式，您可以根据需要选择：

### 方式一：手动 Git Clone

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd donkey-ucenter-backend
   rm -rf .git
   ```

2. **安装依赖**
   ```bash
   go mod download
   ```

3. **初始化配置文件**
   ```bash
   cp conf/config.toml.example conf/config.toml
   # 编辑 conf/config.toml，配置数据库、Redis 等信息
   ```

4. **运行项目**
   ```bash
   go run main.go
   ```

### 方式二：使用 Shell 脚本自动安装

使用 curl 命令一键安装，脚本会自动完成项目克隆、移除 .git 目录、初始化配置文件等操作。

```bash
curl -fsSL https://raw.githubusercontent.com/<your-username>/donkey-ucenter-backend/main/install.sh | bash
```

或者下载脚本后执行：

```bash
curl -fsSL https://raw.githubusercontent.com/<your-username>/donkey-ucenter-backend/main/install.sh -o install.sh
chmod +x install.sh
./install.sh
```

**脚本功能说明：**
- 自动克隆项目到当前目录
- 移除 `.git` 目录（去除版本控制信息）
- 从 `conf/config.toml.example` 复制并初始化 `conf/config.toml`
- 提示用户编辑配置文件

**注意：** 请将 `<repository-url>` 和 `<your-username>` 替换为实际的仓库地址和用户名。

## 配置说明

项目配置文件位于 `conf/config.toml`，主要配置项包括：

- **MySQL 数据库配置**: 数据库连接信息
- **Redis 配置**: 缓存服务配置
- **SMTP 配置**: 邮箱服务配置（用于发送验证邮件）

详细配置请参考 `conf/config.toml.example` 文件。

## 运行

```bash
# 开发环境运行
go run main.go

# 编译运行
go build -o donkey-ucenter-backend
./donkey-ucenter-backend
```

## 许可证

详见 [LICENSE](LICENSE) 文件。
