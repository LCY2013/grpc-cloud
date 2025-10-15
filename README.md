# GRPC Cloud - 完整的 gRPC 云原生解决方案 / Complete gRPC Cloud-Native Solution

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue.svg)](https://golang.org/)
[![gRPC](https://img.shields.io/badge/gRPC-Latest-green.svg)](https://grpc.io/)

## 项目简介 / Project Overview

GRPC Cloud 是一个完整的 gRPC 云原生解决方案，提供了服务注册发现、API 网关、协议验证、日志记录等完整的微服务基础设施组件。该项目旨在简化 gRPC 服务的开发、部署和管理，提供开箱即用的云原生微服务架构支持。

GRPC Cloud is a comprehensive gRPC cloud-native solution that provides complete microservice infrastructure components including service registration and discovery, API gateway, protocol validation, and logging. The project aims to simplify the development, deployment, and management of gRPC services, offering out-of-the-box cloud-native microservice architecture support.

## 核心特性 / Core Features

### 🌐 服务注册与发现 / Service Registration & Discovery
- **自动服务注册**: 服务启动时自动注册到注册中心
- **健康检查**: 内置服务健康检查机制
- **负载均衡**: 支持多种负载均衡策略
- **服务发现**: 动态服务发现和路由

### 🚪 API 网关 / API Gateway
- **HTTP/gRPC 桥接**: 自动将 HTTP 请求转换为 gRPC 调用
- **动态路由**: 基于服务发现的动态路由配置
- **请求转发**: 智能请求转发和负载均衡
- **协议转换**: 支持多种协议间的转换

### ✅ 数据验证 / Data Validation
- **Protobuf 验证**: 基于 protobuf 的消息验证
- **自定义规则**: 支持自定义验证规则
- **实时验证**: 请求级别的实时数据验证
- **错误处理**: 详细的验证错误信息

### 📝 日志系统 / Logging System
- **结构化日志**: 基于 zap 的高性能结构化日志
- **请求追踪**: 完整的请求链路追踪
- **性能监控**: 请求耗时和性能指标
- **可配置级别**: 灵活的日志级别配置

### 🔧 插件系统 / Plugin System
- **中间件支持**: 丰富的 gRPC 中间件
- **可扩展架构**: 易于扩展的插件架构
- **验证插件**: 内置数据验证插件
- **自定义插件**: 支持自定义插件开发

## 项目架构 / Project Architecture

```
grpc-cloud/
├── example/                    # 示例服务 / Example Services
│   ├── echo-service/          # Echo 服务示例 / Echo Service Example
│   └── person-service/        # Person 服务示例 / Person Service Example
├── grpc-gateway/              # API 网关 / API Gateway
│   ├── proto/                 # 协议处理 / Protocol Handling
│   ├── registry/              # 注册中心集成 / Registry Integration
│   └── route/                 # 路由管理 / Route Management
├── logger/                    # 日志系统 / Logging System
├── plugins/                   # 插件系统 / Plugin System
│   └── grpc-cloud-plugin-validate/  # 验证插件 / Validation Plugin
├── proto/                     # Protocol Buffers 定义 / Protocol Buffers Definitions
│   ├── prototpl/              # 模板定义 / Template Definitions
│   ├── gen/                   # 生成的代码 / Generated Code
│   └── buf/                   # Buf 配置 / Buf Configuration
└── registry/                  # 注册中心 / Registry
    ├── registry/              # 注册中心接口 / Registry Interface
    └── grpc-cloud-direct/     # 直接注册实现 / Direct Registry Implementation
```

## 快速开始 / Quick Start

### 环境要求 / Prerequisites

- Go 1.18+ / Go 1.18+
- Protocol Buffers compiler / Protocol Buffers 编译器
- Buf CLI / Buf CLI

### 安装依赖 / Install Dependencies

```bash
# 安装 Buf CLI / Install Buf CLI
go install github.com/bufbuild/buf/cmd/buf@latest

# 安装 Protocol Buffers / Install Protocol Buffers
# macOS / macOS
brew install protobuf

# Ubuntu / Ubuntu
apt-get install protobuf-compiler
```

### 代码生成 / Code Generation

```bash
# 进入 proto 目录 / Enter proto directory
cd proto

# 生成代码 / Generate code
make generate
```

### 启动服务 / Start Services

#### 1. 启动 API 网关 / Start API Gateway

```bash
cd grpc-gateway
go run main.go
```

网关将在 `http://localhost:8080` 启动 / Gateway will start at `http://localhost:8080`

#### 2. 启动示例服务 / Start Example Services

```bash
# 启动 Echo 服务 / Start Echo Service
cd example/echo-service
go run main.go

# 启动 Person 服务 / Start Person Service
cd example/person-service
go run main.go
```

### 测试服务 / Test Services

#### 使用 HTTP 客户端测试 / Test with HTTP Client

```bash
# 测试 Echo 服务 / Test Echo Service
curl -X POST http://localhost:8080/echo/service/v1/example/echo \
  -H "Content-Type: application/json" \
  -d '{"value": "Hello World"}'

# 测试 Person 服务 / Test Person Service
curl -X POST http://localhost:8080/person/service/v1/get \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 25}'
```

#### 使用 gRPC 客户端测试 / Test with gRPC Client

```bash
# 安装 grpcurl / Install grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

# 测试 Echo 服务 / Test Echo Service
grpcurl -plaintext -d '{"value": "Hello World"}' \
  localhost:8888 grpc.echo.service.v1.EchoService/Echo

# 测试 Person 服务 / Test Person Service
grpcurl -plaintext -d '{"name": "John Doe", "age": 25}' \
  localhost:8889 grpc.person.service.v1.PersonService/GetPerson
```

## 详细使用指南 / Detailed Usage Guide

### 服务注册 / Service Registration

服务会自动注册到网关的注册中心。注册信息包括：

Services are automatically registered to the gateway's registry. Registration information includes:

- 服务名称 / Service Name
- 服务 ID / Service ID  
- 服务地址 / Service Address
- 服务端口 / Service Port

```go
// 注册服务 / Register Service
client := grpc_cloud_direct.NewDirectRegistryClient("http://127.0.0.1:8080")
err := client.Register("echo-service", "echo-service", "127.0.0.1", 8888)
```

### 数据验证 / Data Validation

项目支持两种验证方式：

The project supports two validation methods:

#### 1. 原生 Protobuf 验证 / Native Protobuf Validation

```go
// 使用原生验证拦截器 / Use native validation interceptor
grpc.ChainUnaryInterceptor(
    grpc_cloud_plugin_validate.WithValidateServerInterceptor,
)
```

#### 2. Buf 验证 / Buf Validation

```go
// 使用 Buf 验证拦截器 / Use Buf validation interceptor
grpc.ChainUnaryInterceptor(
    grpc_cloud_plugin_validate.WithValidatePbServerInterceptor,
)
```

### 自定义验证规则 / Custom Validation Rules

在 `.proto` 文件中定义验证规则：

Define validation rules in `.proto` files:

```protobuf
message StringMessage {
  string value = 1 [
    (validate.rules).string = {min_len: 5}
  ];
}

message PersonMessage {
  string name = 1 [(buf.validate.field).string.min_len = 5];
  int32 age = 2 [(buf.validate.field).int32.gt = 20];
}
```

### 日志配置 / Logging Configuration

项目使用 zap 作为日志库，支持结构化日志：

The project uses zap as the logging library with structured logging support:

```go
import "github.com/LCY2013/grpc-cloud/logger"

// 使用全局日志实例 / Use global logger instance
logger.Log.Info("Service started")
logger.Log.Error("Service error", err)
```

## API 文档 / API Documentation

### Echo Service API

| 方法 / Method | HTTP 路径 / HTTP Path | 描述 / Description |
|--------------|---------------------|-------------------|
| Echo | POST `/echo/service/v1/example/echo` | 回显消息 / Echo message |
| EchoCustomer | POST `/echo/service/v1/example/echo/c` | 客户回显 / Customer echo |

### Person Service API

| 方法 / Method | HTTP 路径 / HTTP Path | 描述 / Description |
|--------------|---------------------|-------------------|
| GetPerson | POST `/person/service/v1/get` | 获取人员信息 / Get person info |

### 请求/响应格式 / Request/Response Format

#### Echo Service

**请求 / Request:**
```json
{
  "value": "Hello World"
}
```

**响应 / Response:**
```json
{
  "value": "hello, Hello World"
}
```

#### Person Service

**请求 / Request:**
```json
{
  "name": "John Doe",
  "age": 25
}
```

**响应 / Response:**
```json
{
  "name": "hello, John Doe",
  "age": 25
}
```

## 配置选项 / Configuration Options

### 网关配置 / Gateway Configuration

```go
// 初始化注册中心 / Initialize registry
registry.Init()

// 初始化路由 / Initialize router
httprouter.Init()

// 启动 HTTP 服务器 / Start HTTP server
http.ListenAndServe(":8080", nil)
```

### 服务配置 / Service Configuration

```go
// 创建 gRPC 服务器 / Create gRPC server
svr := grpc.NewServer(
    grpc.ChainUnaryInterceptor(
        grpc_cloud_plugin_validate.WithValidateServerInterceptor,
    ),
    grpc.ChainStreamInterceptor(),
)

// 注册服务 / Register service
v1.RegisterEchoServiceServer(svr, &EchoServiceProc{})

// 启用反射 / Enable reflection
reflection.Register(svr)
```

## 开发指南 / Development Guide

### 添加新服务 / Adding New Services

1. **创建 Proto 文件 / Create Proto File**

```protobuf
syntax = "proto3";
package grpc.yourapp.service.v1;
option go_package = "grpc/yourapp/service/v1";

import "google/api/annotations.proto";
import "validate/validate.proto";

message YourMessage {
  string field = 1 [(validate.rules).string = {min_len: 1}];
}

service YourService {
  rpc YourMethod(YourMessage) returns (YourMessage) {
    option (google.api.http) = {
      post: "/yourapp/service/v1/yourmethod"
      body: "*"
    };
  }
}
```

2. **生成代码 / Generate Code**

```bash
cd proto
make generate
```

3. **实现服务 / Implement Service**

```go
type YourServiceProc struct {
    v1.YourServiceServer
}

func (y *YourServiceProc) YourMethod(ctx context.Context, req *v1.YourMessage) (*v1.YourMessage, error) {
    // 实现业务逻辑 / Implement business logic
    return &v1.YourMessage{
        Field: "processed: " + req.Field,
    }, nil
}
```

4. **启动服务 / Start Service**

```go
func main() {
    svr := grpc.NewServer(
        grpc.ChainUnaryInterceptor(
            grpc_cloud_plugin_validate.WithValidateServerInterceptor,
        ),
    )
    
    v1.RegisterYourServiceServer(svr, &YourServiceProc{})
    reflection.Register(svr)
    
    // 注册到网关 / Register to gateway
    go func() {
        time.Sleep(5 * time.Second)
        client := grpc_cloud_direct.NewDirectRegistryClient("http://127.0.0.1:8080")
        client.Register("your-service", "your-service", "127.0.0.1", 8888)
    }()
    
    l, _ := net.Listen("tcp", "127.0.0.1:8888")
    svr.Serve(l)
}
```

### 自定义插件开发 / Custom Plugin Development

```go
package yourplugin

import (
    "context"
    "google.golang.org/grpc"
)

func WithYourServerInterceptor() grpc.UnaryServerInterceptor {
    return func(ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler) (interface{}, error) {
        
        // 前置处理 / Pre-processing
        // ... your logic here ...
        
        // 调用处理器 / Call handler
        resp, err := handler(ctx, req)
        
        // 后置处理 / Post-processing
        // ... your logic here ...
        
        return resp, err
    }
}
```

## 部署指南 / Deployment Guide

### Docker 部署 / Docker Deployment

```dockerfile
# Dockerfile 示例 / Dockerfile Example
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o gateway ./grpc-gateway/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/gateway .
EXPOSE 8080
CMD ["./gateway"]
```

### Kubernetes 部署 / Kubernetes Deployment

```yaml
# k8s-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: grpc-cloud-gateway
spec:
  replicas: 3
  selector:
    matchLabels:
      app: grpc-cloud-gateway
  template:
    metadata:
      labels:
        app: grpc-cloud-gateway
    spec:
      containers:
      - name: gateway
        image: grpc-cloud/gateway:latest
        ports:
        - containerPort: 8080
        env:
        - name: GATEWAY_PORT
          value: "8080"
---
apiVersion: v1
kind: Service
metadata:
  name: grpc-cloud-gateway-service
spec:
  selector:
    app: grpc-cloud-gateway
  ports:
  - port: 8080
    targetPort: 8080
  type: LoadBalancer
```

## 性能优化 / Performance Optimization

### 连接池配置 / Connection Pool Configuration

```go
// 配置 gRPC 客户端连接 / Configure gRPC client connection
conn, err := grpc.Dial("localhost:8888",
    grpc.WithInsecure(),
    grpc.WithKeepaliveParams(keepalive.ClientParameters{
        Time:                10 * time.Second,
        Timeout:             3 * time.Second,
        PermitWithoutStream: true,
    }),
)
```

### 日志级别优化 / Log Level Optimization

```go
// 生产环境日志配置 / Production logging configuration
config := zap.NewProductionConfig()
config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
logger, _ := config.Build()
```

## 故障排除 / Troubleshooting

### 常见问题 / Common Issues

#### 1. 服务注册失败 / Service Registration Failed

**问题 / Problem:** 服务无法注册到网关 / Service cannot register to gateway

**解决方案 / Solution:**
- 检查网关是否运行 / Check if gateway is running
- 验证注册地址和端口 / Verify registration address and port
- 检查网络连接 / Check network connectivity

#### 2. 验证失败 / Validation Failed

**问题 / Problem:** 请求数据验证失败 / Request data validation failed

**解决方案 / Solution:**
- 检查 proto 文件中的验证规则 / Check validation rules in proto files
- 验证请求数据格式 / Verify request data format
- 查看详细错误信息 / Check detailed error messages

#### 3. 路由不匹配 / Route Not Found

**问题 / Problem:** HTTP 请求无法找到对应的 gRPC 服务 / HTTP request cannot find corresponding gRPC service

**解决方案 / Solution:**
- 检查 proto 文件中的 HTTP 注解 / Check HTTP annotations in proto files
- 确认服务已正确注册 / Confirm service is properly registered
- 验证路由配置 / Verify route configuration

## 贡献指南 / Contributing

我们欢迎社区贡献！请遵循以下步骤：

We welcome community contributions! Please follow these steps:

1. **Fork 项目 / Fork the project**
2. **创建特性分支 / Create feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **提交更改 / Commit changes**
   ```bash
   git commit -m 'Add amazing feature'
   ```
4. **推送分支 / Push branch**
   ```bash
   git push origin feature/amazing-feature
   ```
5. **创建 Pull Request / Create Pull Request**

### 代码规范 / Code Standards

- 使用 Go 标准格式 / Use Go standard formatting
- 添加适当的注释 / Add appropriate comments
- 编写单元测试 / Write unit tests
- 更新文档 / Update documentation

## 许可证 / License

本项目采用 Apache License 2.0 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## 联系方式 / Contact

- **项目主页 / Project Homepage:** [https://github.com/LCY2013/grpc-cloud](https://github.com/LCY2013/grpc-cloud)
- **问题反馈 / Issue Tracker:** [https://github.com/LCY2013/grpc-cloud/issues](https://github.com/LCY2013/grpc-cloud/issues)
- **邮箱 / Email:** none@example.com

## 致谢 / Acknowledgments

感谢以下开源项目：

Thanks to the following open source projects:

- [gRPC](https://grpc.io/) - 高性能 RPC 框架 / High-performance RPC framework
- [Protocol Buffers](https://developers.google.com/protocol-buffers) - 序列化框架 / Serialization framework
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) - gRPC 到 HTTP 的代理 / gRPC to HTTP proxy
- [zap](https://github.com/uber-go/zap) - 结构化日志库 / Structured logging library
- [buf](https://buf.build/) - Protocol Buffers 工具链 / Protocol Buffers toolchain

---

**Happy Coding! / 编码愉快！** 🚀