# newb-cli

爱航·Golang快速开发模板CLI工具

## 安装

```bash
go install github.com/lovehang/newb-cli@latest
```

## 命令说明

### 查看版本
```bash
newb version
```

### 创建新项目

基础用法:
```bash
newb new myapp
```

高级选项:
```bash
# 指定模板仓库
newb new myapp --template https://github.com/your/repo

# 指定分支
newb new myapp --branch develop

# 强制覆盖已存在目录
newb new myapp --force

# 组合使用
newb new myapp -t https://github.com/your/repo -b develop -f
```

### 命令参数

| 参数 | 简写 | 说明 | 默认值 |
|------|------|------|--------|
| --template | -t | 模板仓库地址 | https://github.com/lxhanghub/newb |
| --branch | -b | 模板仓库分支 | main |
| --force | -f | 强制覆盖已存在目录 | false |
| --config | 无 | 指定配置文件路径 | $HOME/.newb.yaml |

## 项目结构

使用newb创建的项目将包含以下结构:

```
myapp/
  ├── cmd/        # 命令行入口
  ├── internal/   # 内部包
  ├── pkg/        # 公共包
  ├── go.mod      # 依赖管理
  └── main.go     # 主程序入口
```

## 配置文件

默认配置文件位置: `$HOME/.newb.yaml`

配置示例:
```yaml
template:
  default: https://github.com/lxhanghub/newb
  branch: main
```

## 本地开发

```bash
# 克隆项目
git clone https://github.com/lovehang/newb-cli.git
cd newb-cli

# 方式1: 直接运行
go run main.go

# 方式2: 编译后运行
go build -o newb main.go
./newb

# 测试命令
./newb version
./newb new myapp
```

## 许可证

MIT License
