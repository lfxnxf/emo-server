FROM xuefeng6329/meta-sports

ADD api-gateway /app/api-gateway
COPY ./config.toml /config.toml
COPY ./pem /pem

#执行命令
ENTRYPOINT ["/app/api-gateway", "-config", "/config.toml"]

#######
#基础镜像 xuefeng6329/meta-sports
#FROM ubuntu:18.04
#
#RUN  sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list
#RUN  sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list
#RUN  apt-get clean
#RUN apt-get -qq update \
#    && apt-get -qq install -y --no-install-recommends ca-certificates curl

#FROM golang:alpine AS builder
#
#LABEL stage=gobuilder
#
#ENV CGO_ENABLED 0
#ENV GOPROXY https://goproxy.cn,direct
#
#RUN apk --no-cache add tzdata
#
#WORKDIR /build
#
#COPY . .
#RUN go build -o /app/api-gateway ./app
#COPY ./config.toml /app
#COPY ./pem /app/pem
#
#
#FROM alpine
#
#RUN apk --no-cache add ca-certificates
#COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
#ENV TZ Asia/Shanghai
#
#WORKDIR /app
#COPY --from=builder /app/api-gateway /app/api-gateway
#COPY --from=builder /app/config.toml /config.toml
#COPY --from=builder /app/pem /pem
#
#CMD ["/app/api-gateway", "-config", "/config.toml"]
