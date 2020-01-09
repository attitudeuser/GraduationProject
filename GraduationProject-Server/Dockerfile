FROM golang:latest

LABEL maintainer="xuthus5@gmail.com" version=1.0.alpha

WORKDIR /Project

RUN go env -w GOPROXY=https://goproxy.cn,direct

EXPOSE 5000

ENTRYPOINT ["go","run","main.go"]
