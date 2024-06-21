# 构建Docker 镜像模版
# 构建阶段
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o phone main.go && apk --no-cache add upx && upx -9 phone

# 最终镜像
FROM alpine
WORKDIR /app
COPY --from=builder /app/phone .
ENTRYPOINT ["./phone"]
