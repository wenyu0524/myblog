FROM golang:1.26-bookworm AS build

WORKDIR /src/blog-rpc
ENV GOPROXY=https://goproxy.cn,direct
COPY blog-rpc/go.mod blog-rpc/go.sum ./
RUN go mod download
COPY blog-rpc/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/blog-rpc .

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/blog-rpc ./blog-rpc
COPY blog-rpc/etc ./etc
ENTRYPOINT ["/app/blog-rpc", "-f", "/app/etc/blog.yaml"]
