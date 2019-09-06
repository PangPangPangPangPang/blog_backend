FROM golang:latest

RUN go env

# 打包是定义密钥用于执行更新
ENV BBE_SECRET_KEY
# 打包为release
ENV GIN_MODE release

ENV GO111MODULE on

ENV GOPATH /go:/blog_backend_src

RUN mkdir -p /blog_backend_src

COPY . /blog_backend_src
WORKDIR /blog_backend_src

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./blog_backend_src"]
