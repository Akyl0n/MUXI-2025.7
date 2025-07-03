# 使用官方Go镜像作为构建环境和运行环境
FROM golang:1.24

WORKDIR /app
COPY . .

RUN go mod init go-book-app || true
RUN go mod tidy
RUN go build -o app t1.go

CMD ["./app"]