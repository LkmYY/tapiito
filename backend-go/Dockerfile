# 多阶段构建：编译产物 + 最小运行镜像
# Go 版本需 >= go.mod 声明（当前 1.26.3），用 ARG 便于升级
ARG GO_VERSION=1.26
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /src
ENV CGO_ENABLED=0 GOOS=linux
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api && \
    go build -trimpath -ldflags="-s -w" -o /out/worker ./cmd/worker && \
    go build -trimpath -ldflags="-s -w" -o /out/bloom-builder ./cmd/bloom-builder

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata && \
    adduser -D -u 10001 app
WORKDIR /app
COPY --from=builder /out/api /out/worker /out/bloom-builder /app/
COPY migrations /app/migrations
USER app
EXPOSE 8080 8081
ENTRYPOINT ["/app/api"]
