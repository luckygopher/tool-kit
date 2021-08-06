# 多阶段构建：编译阶段
FROM golang:1.13.5-alpine3.10 AS builder
# 指定工作目录，相当于 cd 命令
WORKDIR /build
# 添加一个无密码的用户 app-runner 用来启动应用
RUN adduser -u 10001 -D app-runner
# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"
# 安装依赖包，可以避免重复下载，加快构建
COPY go.mod .
COPY go.sum .
RUN go mod download
# 拷贝代码到容器中
COPY . .
# 编译代码为二进制可执行文件app
RUN go build -a -o app .

# 最终镜像
FROM alpine3.10 AS final
WORKDIR /app
# 这里把编译阶段编译好的二进制文件、配置等拷贝过来
COPY --from=builder /build/app /app/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# 使用 app-runner 用户启动应用
USER app-runner
ENTRYPOINT ["/app/app"]

