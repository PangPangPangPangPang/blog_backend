FROM golang:latest

# 打包是定义密钥用于执行更新
ENV BBE_SECRET_KEY

WORKDIR /root/blog_backend
COPY . /root/blog_backend

RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/contrib/static

RUN go build .

EXPOSE 8080

ENTRYPOINT ["./blog_backend"]
