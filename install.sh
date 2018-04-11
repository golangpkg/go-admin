#!/bin/sh


echo "deb http://mirrors.aliyun.com/debian/ stretch main non-free contrib" > /etc/apt/sources.list && \
echo "deb http://mirrors.aliyun.com/debian/ stretch-proposed-updates main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.aliyun.com/debian/ stretch main non-free contrib" >> /etc/apt/sources.list && \
echo "deb-src http://mirrors.aliyun.com/debian/ stretch-proposed-updates main non-free contrib" >> /etc/apt/sources.list

apt-get update && apt-get install -y libsqlite3-dev

env GOOS=linux GOARCH=amd64 CGO_ENABLED=1

cd /go/src/github.com/golangpkg/go-admin
rm -f go-admin
go build -ldflags "-linkmode external -extldflags -static" -o go-admin main.go

echo "install finish."