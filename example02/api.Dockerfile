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
COPY api/*.go ./api/
COPY demo-api/*.go ./

# 编译代码
RUN go build -o /app/demo-api

EXPOSE 9000
EXPOSE 9001

CMD [ "/app/demo-api"]