FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR $GOPATH/src/github.com/lius-new/liusnew-blog-backend-server
COPY . $GOPATH/src/github.com/lius-new/liusnew-blog-backend-server

RUN go build -o liusnew-blog-backend-server .

FROM debian:stretch-slim

COPY --from=builder $GOPATH/src/github.com/lius-new/liusnew-blog-backend-server/liusnew-blog-backend-server /

ENTRYPOINT ["./liusnew-blog-backend-server"]