FROM golang:latest

RUN go env

# 打包是定义密钥用于执行更新
ENV BBE_SECRET_KEY jkadsfvbaiojwaerklaw
# 打包为release
ENV GIN_MODE release

RUN mkdir -p /go/src/github.com/PangPangPangPangPang
COPY . /go/src/github.com/PangPangPangPangPang/blog_backend
WORKDIR /go/src/github.com/PangPangPangPangPang/blog_backend

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./blog_backend"]
