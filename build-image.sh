#!/bin/bash

set -xe

IMAGE=${1:?"docker image name required"}

echo
echo "Testing and compiling..."
echo

docker run -v `pwd`:/app -w /app                     golang:1.8.3 go test
docker run -v `pwd`:/app -w /app --env CGO_ENABLED=0 golang:1.8.3 go build -o server -a -tags netgo -ldflags '-w'

echo
echo "Building docker image $IMAGE ..."
echo

cat >Dockerfile <<EOF
FROM scratch
COPY server server
EOF
docker build -t $IMAGE .

echo
echo "Cleaning up..."
echo

rm server
rm Dockerfile

echo
echo "Built docker image $IMAGE:"
echo

docker image inspect $IMAGE
