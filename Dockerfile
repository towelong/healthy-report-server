FROM golang:1.16.9 AS builder
WORKDIR /go/src
COPY . .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.io,direct \
    && CGO_ENABLED=0 go build -o App main.go

FROM alpine AS final
WORKDIR /app
COPY --from=builder /go/src/App /app/App
COPY --from=builder /go/src/.env.production /app/App/.env.production
RUN chmod a+xr -R /app/App
EXPOSE 8016
ENTRYPOINT ["/app/App","-conf", "/app/App/.env.production"]