FROM golang:1.26-bookworm AS build

WORKDIR /src/user-rpc
ENV GOPROXY=https://goproxy.cn,direct
COPY user-rpc/go.mod user-rpc/go.sum ./
RUN go mod download
COPY user-rpc/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/user-rpc .

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/user-rpc ./user-rpc
COPY user-rpc/etc ./etc
ENTRYPOINT ["/app/user-rpc", "-f", "/app/etc/user.yaml"]
