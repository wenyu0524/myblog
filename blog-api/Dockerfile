# blog-api Dockerfile - 多阶段构建
# 构建上下文需为项目根目录（docker-compose 中 context: .）
# blog-api 的 go.mod 有 replace 指令指向 ../user-rpc 和 ../blog-rpc，
# 因此必须同时复制这两个模块的代码

# =========================
# 阶段一：编译 Go 二进制
# =========================
FROM golang:1.26 AS builder

ENV GOPROXY=https://goproxy.cn,direct
ENV GOSUMDB=off

WORKDIR /build

# 复制全部三个模块（blog-api 通过 replace 引用 user-rpc 和 blog-rpc）
COPY user-rpc/ ./user-rpc/
COPY blog-rpc/ ./blog-rpc/
COPY blog-api/ ./blog-api/

WORKDIR /build/blog-api

# 静态编译
RUN CGO_ENABLED=0 GOOS=linux go build -o /blog-api .

# =========================
# 阶段二：最小化运行镜像
# =========================
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /blog-api /app/blog-api
COPY blog-api/etc/blog-api.yaml /app/etc/blog-api.yaml

WORKDIR /app

ENTRYPOINT ["/app/blog-api", "-f", "etc/blog-api.yaml"]