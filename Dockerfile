FROM golang:latest

MAINTAINER Goblog "2387805574@qq.com"

WORKDIR /Project/doit

#COPY go.mod go.sum ../
#RUN go mod download

COPY . /Project

WORKDIR /Project

ENTRYPOINT ["./doit/doit"]
