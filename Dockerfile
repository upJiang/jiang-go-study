# 使用官方 Golang 镜像作为基础镜像
FROM golang:1.22.5

# 设置工作目录
WORKDIR /app

# 复制所有文件到容器中
COPY . .

# 下载并安装依赖包
RUN go mod tidy

# 编译 Go 程序
RUN go build -o jiang-study-go

# 声明服务端口
EXPOSE 8080

# 启动 Go 应用
CMD ["./jiang-study-go"]
