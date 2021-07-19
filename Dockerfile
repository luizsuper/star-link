FROM golang:buster
ENV GOPROXY https://goproxy.cn/
WORKDIR /go/src/app
COPY . .
