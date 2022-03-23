FROM golang:1.16.9 AS builder
WORKDIR /go/src
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn && CGO_ENABLED=0 go build -o App main.go

FROM alpine AS final
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
  apk --update add tzdata && \
  cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
  echo "Asia/Shanghai" > /etc/timezone && \
  apk del tzdata && \
  rm -rf /var/cache/apk/*

WORKDIR /app
COPY --from=builder /go/src/App /app/App
COPY --from=builder /go/src/.env.production /app/.env.production
RUN chmod a+xr -R /app/App
EXPOSE 8016
ENTRYPOINT ["/app/App","-conf",".env.production"]