#!/bin/bash
docker build --rm -t api-gateway .

host=$(ip addr | awk '/^[0-9]+: / {}; /inet.*global.*eth/ {print gensub(/(.*)\/(.*)/, "\\1", "g", $2)}')

docker image prune -f

docker stop api-gateway-1
docker rm api-gateway-1
docker run -d --restart always --name api-gateway-1 \
-p 10000:10000 \
-v /work/logs/api-gateway:/logs \
-v /work/api/images:/images \
-v /etc/localtime:/etc/localtime:ro \
-e FINAL_HOST=${host} api-gateway

docker stop api-gateway-2
docker rm api-gateway-2
docker run -d --restart always --name api-gateway-2 \
-p 10001:10000 \
-v /work/logs/api-gateway:/logs \
-v /work/api/images:/images \
-v /etc/localtime:/etc/localtime:ro \
-e FINAL_HOST=${host} api-gateway
