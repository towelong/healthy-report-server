FROM golang:1.16.9 AS builder
WORKDIR /go/src
COPY . .
RUN GOPROXY=https://goproxy.cn && go build -o App main.go

FROM alpine AS final
WORKDIR /app
COPY --from=builder /go/src/App /app/App
COPY --from=builder /go/src/.env.production /app/.env.production
RUN chmod a+xr -R /app/App
EXPOSE 8016
ENTRYPOINT ["/app/App","-conf", "/app/.env.production"]