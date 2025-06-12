# workit-cli

爱航·Golang快速开发模板CLI工具

## 安装

### 方式一：go install 安装（推荐）

```bash
# 设置 GOPROXY 环境变量（解决可能的网络问题）
go env -w GOPROXY=https://goproxy.cn,direct

# 安装最新版本
go install github.com/lxhanghub/workit-cli/cmd/workit@latest

# 或安装指定版本
go install github.com/lxhanghub/workit-cli/cmd/workit@v1.0.8
```

### 方式二：手动编译安装

如果 go install 安装失败，可以尝试手动编译安装：

```bash
# 克隆仓库
git clone https://github.com/lxhanghub/workit-cli.git
cd workit-cli

# 编译
go build -o workit cmd/workit/main.go

# 将编译好的程序移动到 GOPATH/bin 目录
mv newb $GOPATH/bin/
```

### 常见问题排查

1. 确认 Go 环境配置：

```bash
# 检查 Go 版本（需要 1.16+）
go version

# 检查 GOPATH 和 GOPROXY
go env GOPATH GOPROXY
```

2. 如果出现网络问题，尝试：

```bash
# 清理模块缓存
go clean -modcache

# 重新下载依赖
go mod download
```

## 命令说明

### 查看版本

```bash
workit version
```

### 创建新项目

基础用法:

```bash
workit new myapp
```

高级选项:

```bash
# 指定模板仓库
workit new myapp --template https://github.com/your/repo  

# 指定分支
workit new myapp --branch cli-echo

# 强制覆盖已存在目录
workit new myapp --force

# 组合使用
workit new myapp -t https://github.com/your/repo -b develop -f
```

### 命令参数


| 参数       | 简写 | 说明               | 默认值                              |
| ---------- | ---- | ------------------ | ----------------------------------- |
| --template | -t   | 模板仓库地址       | https://github.com/lxhanghub/workit |
| --branch   | -b   | 模板仓库分支       | cli-template                        |
| --force    | -f   | 强制覆盖已存在目录 | false                               |
| --config   | 无   | 指定配置文件路径   | $HOME/.newb.yaml                    |

## 项目结构

使用workit创建的项目将包含以下结构:

```
  myapp/
  ├── api/        # api规范
  ├── build/      # 构建文件
  ├── cmd/        # 命令行入口
  |    └── service1/main.go     # 主程序入口
  ├── deployment/ # 部署文件
  ├── internal/   # 内部包
  |   └── service1            # 微服务   
  |       └── application     # 应用层
  |       └── domain          # 领域层
  |       └── grpcapi         # rpc 接口层
  |       └── infrastructure  # 基础设施层 
  |       └── webapi          # webapi接口层
  ├── pkg/        # 公共包
  ├── go.mod      # 依赖管理
```

## 配置文件

默认配置文件位置: `$HOME/.workit.yaml`

配置示例:

```yaml
template:
  default: https://github.com/lxhanghub/workit
  branch: cli-template
```

## 许可证

MIT License
