# user-rpc Dockerfile - 多阶段构建
# 构建上下文需为项目根目录（docker-compose 中 context: .）

# =========================
# 阶段一：编译 Go 二进制
# =========================
FROM golang:1.26 AS builder

# 设置 Go 模块代理和校验配置（适配国内网络环境）
ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

WORKDIR /build

# 仅复制 user-rpc 模块代码（user-rpc 无外部 replace 依赖）
COPY user-rpc/ ./user-rpc/

WORKDIR /build/user-rpc

# 静态编译，CGO_ENABLED=0 确保 pgx 纯 Go 驱动不依赖 C 库
RUN CGO_ENABLED=0 GOOS=linux go build -o /user-rpc .

# =========================
# 阶段二：最小化运行镜像
# =========================
FROM alpine:latest

# 安装 CA 证书（gRPC TLS 通信需要）和时区数据
RUN apk --no-cache add ca-certificates tzdata

# 从构建阶段复制二进制和配置文件
COPY --from=builder /user-rpc /app/user-rpc
COPY user-rpc/etc/user.yaml /app/etc/user.yaml

WORKDIR /app

ENTRYPOINT ["/app/user-rpc", "-f", "etc/user.yaml"]