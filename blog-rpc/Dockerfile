# blog-rpc Dockerfile - 多阶段构建
# 构建上下文需为项目根目录（docker-compose 中 context: .）

# =========================
# 阶段一：编译 Go 二进制
# =========================
FROM golang:1.26 AS builder

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

WORKDIR /build

# 仅复制 blog-rpc 模块代码（blog-rpc 无外部 replace 依赖）
COPY blog-rpc/ ./blog-rpc/

WORKDIR /build/blog-rpc

# 静态编译
RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-rpc .

# =========================
# 阶段二：最小化运行镜像
# =========================
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /blog-rpc /app/blog-rpc
COPY blog-rpc/etc/blog.yaml /app/etc/blog.yaml

WORKDIR /app

ENTRYPOINT ["/app/blog-rpc", "-f", "etc/blog.yaml"]