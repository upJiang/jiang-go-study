# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.20

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到 /app 目录
COPY . .

# 下载依赖
RUN go mod tidy

# 编译应用
RUN go build -o my-gin-app .

# 启动应用
CMD ["./my-gin-app"]
