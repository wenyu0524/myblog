FROM golang:1.26-bookworm AS build

WORKDIR /src
COPY go.work go.work.sum ./
COPY blog-rpc/ ./blog-rpc/
COPY myblog-api/ ./myblog-api/
COPY user-rpc/ ./user-rpc/
RUN CGO_ENABLED=0 GOOS=linux go build -C myblog-api -o /out/myblog-api .

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/myblog-api ./myblog-api
COPY myblog-api/etc ./etc
ENTRYPOINT ["/app/myblog-api", "-f", "/app/etc/myblog-api.yaml"]
