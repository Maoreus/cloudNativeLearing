FROM golang:1.18.3-alpine3.16 AS build
WORKDIR /workdir
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -o httpserver main.go

#原来没写分阶段构建，现在加上
FROM busybox
COPY --from=build --chown=user /httpserver/workdir /httpserver/workdir
EXPOSE 9999
WORKDIR /build
ENTRYPOINT ["./httpserver"]

#docker build -t httpserver:01
#docker push Maoreus/httpserver:01



