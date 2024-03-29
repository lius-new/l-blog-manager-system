FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR $GOPATH/src/github.com/lius-new/liusnew-blog-backend-server
COPY . .

RUN go build -o liusnew-blog-backend-server .

FROM debian:stretch-slim

# /go/src/github.com/lius-new/liusnew-blog-backend-server/liusnew-blog-backend-server
# 为什么不是$GOPATH/src/github.com, 因为一直报错,发现GOPATH就是/go
COPY --from=builder /go/src/github.com/lius-new/liusnew-blog-backend-server/liusnew-blog-backend-server /
COPY --from=builder /go/src/github.com/lius-new/liusnew-blog-backend-server/.env.example /.env

ENTRYPOINT ["./liusnew-blog-backend-server"]