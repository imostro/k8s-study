FROM golang:1.19-alpine

# 指定工作目录
WORKDIR /app

# 复制依赖
COPY go.mod ./
COPY go.sum ./

# 设置go代理
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
#RUN go env -w GOPROXY=https://goproxy.cn,direct

# 运行命令下载依赖
RUN go mod download

# 复制源码
COPY *.go ./

# 编译代码
RUN go build -o /app/example_server

EXPOSE 9000

CMD [ "/app/example_server" ]