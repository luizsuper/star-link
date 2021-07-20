FROM golang:buster
ENV GOPROXY https://goproxy.cn/
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main