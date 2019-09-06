FROM golang:latest

RUN go env

# 打包是定义密钥用于执行更新
ENV BBE_SECRET_KEY
# 打包为release
ENV GIN_MODE release

ENV GO111MODULE on

ENV GOPATH /go:/root/blog_backend

RUN mkdir -p /root/blog_backend

COPY . /root/blog_backend
WORKDIR /root/blog_backend

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./blog_backend"]
