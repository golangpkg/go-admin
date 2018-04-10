#!/bin/sh

git pull

BUILD_DATE=`date +%Y-%m-%d:%H:%M:%S`
echo `date +%Y-%m-%d:%H:%M:%S` > version

base_url=
docker_url=freewebsys/go-admin


#docker run --rm -v `pwd`:/go/src/github.com/freewebsys/go-admin -v `pwd`:/go -i golang:stretch \
#    /go/src/github.com/freewebsys/go-admin/install.sh

docker build -t ${docker_url} .

echo "docker login -u  ${base_url}"
#docker push ${docker_url}

echo "build & push finish ."
echo "##############  $BUILD_DATE  ##############"
