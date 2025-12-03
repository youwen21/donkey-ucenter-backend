#!/bin/bash

# donkey-ucenter-backend 自动安装脚本
# 功能：克隆项目、移除 .git、初始化配置文件

set -e

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 默认仓库地址（请根据实际情况修改）
# 支持通过命令行参数或环境变量 REPO_URL 传递仓库地址
REPO_URL="${1:-${REPO_URL:-https://github.com/your-username/donkey-ucenter-backend.git}}"
PROJECT_NAME="donkey-ucenter-backend"
CONFIG_EXAMPLE="conf/config.toml.example"
CONFIG_FILE="conf/config.toml"

echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}donkey-ucenter-backend 自动安装脚本${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo -e "${YELLOW}使用方法:${NC}"
echo "  ./install.sh [仓库地址]"
echo "  或设置环境变量: REPO_URL=<仓库地址> ./install.sh"
echo ""

# 检查 git 是否安装
if ! command -v git &> /dev/null; then
    echo -e "${RED}错误: 未检测到 git，请先安装 git${NC}"
    exit 1
fi

# 如果项目目录已存在，询问是否删除
if [ -d "$PROJECT_NAME" ]; then
    echo -e "${YELLOW}警告: 目录 $PROJECT_NAME 已存在${NC}"
    read -p "是否删除并重新安装? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        echo -e "${YELLOW}正在删除旧目录...${NC}"
        rm -rf "$PROJECT_NAME"
    else
        echo -e "${RED}安装已取消${NC}"
        exit 1
    fi
fi

# 克隆项目
echo -e "${GREEN}[1/3] 正在克隆项目...${NC}"
echo "仓库地址: $REPO_URL"
git clone "$REPO_URL" "$PROJECT_NAME" || {
    echo -e "${RED}错误: 克隆项目失败，请检查仓库地址和网络连接${NC}"
    exit 1
}
echo -e "${GREEN}✓ 项目克隆成功${NC}"
echo ""

# 进入项目目录
cd "$PROJECT_NAME"

# 移除 .git 目录
echo -e "${GREEN}[2/3] 正在移除 .git 目录...${NC}"
if [ -d ".git" ]; then
    rm -rf .git
    echo -e "${GREEN}✓ .git 目录已移除${NC}"
else
    echo -e "${YELLOW}警告: .git 目录不存在，跳过${NC}"
fi
echo ""

# 初始化配置文件
echo -e "${GREEN}[3/3] 正在初始化配置文件...${NC}"
if [ -f "$CONFIG_EXAMPLE" ]; then
    if [ -f "$CONFIG_FILE" ]; then
        echo -e "${YELLOW}警告: $CONFIG_FILE 已存在，跳过复制${NC}"
    else
        cp "$CONFIG_EXAMPLE" "$CONFIG_FILE"
        echo -e "${GREEN}✓ 配置文件已初始化: $CONFIG_FILE${NC}"
    fi
else
    echo -e "${RED}错误: 未找到配置文件模板 $CONFIG_EXAMPLE${NC}"
    exit 1
fi
echo ""

# 安装 Go 依赖
if command -v go &> /dev/null; then
    echo -e "${GREEN}正在安装 Go 依赖...${NC}"
    go mod download || {
        echo -e "${YELLOW}警告: Go 依赖安装失败，请手动执行 'go mod download'${NC}"
    }
    echo -e "${GREEN}✓ Go 依赖安装完成${NC}"
    echo ""
else
    echo -e "${YELLOW}警告: 未检测到 Go，请手动安装依赖: go mod download${NC}"
    echo ""
fi

# 完成提示
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN}安装完成！${NC}"
echo -e "${GREEN}========================================${NC}"
echo ""
echo -e "${YELLOW}下一步操作：${NC}"
echo "1. 进入项目目录: cd $PROJECT_NAME"
echo "2. 编辑配置文件: $CONFIG_FILE"
echo "3. 运行项目: go run main.go"
echo ""
echo -e "${YELLOW}配置文件位置: ${NC}$(pwd)/$CONFIG_FILE"
echo ""

