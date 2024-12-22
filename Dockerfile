# 使用官方的 Golang 镜像作为基础镜像
FROM ubuntu:22.04 AS builder

# 更新包列表并安装必要的依赖
RUN apt-get update && apt-get install -y build-essential wget && rm -rf /var/lib/apt/lists/*

# 下载并解压 Go 二进制发行版
ENV GO_VERSION=1.20
ENV GO_DOWNLOAD_URL=https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz
RUN wget ${GO_DOWNLOAD_URL} && \
    tar -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz && \
    mv go /usr/local

ENV PATH=/usr/local/go/bin:${PATH}

RUN go version

ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 将项目的 Go 源码复制到 Docker 容器中
COPY . /app

# 构建你的 Go 应用
#RUN go build -o gin-demo ./cmd/main.go #使用Makefile替代
RUN make clean && make

# 使用一个更小的镜像来运行你的应用
FROM alpine:latest
RUN apk add --no-cache bash curl postgresql-client
WORKDIR /app

# 从构建阶段复制可执行文件到运行阶段
# 假设 build.sh 生成的可执行文件在当前目录下（即 /app/gin-demo）
COPY --from=builder /app/gin-demo .
COPY --from=builder /app/resources ./resources

#CMD ["/app/gin-demo", "-c", "/app/resources/config.yaml"]