FROM golang:alpine as builder
WORKDIR /build
RUN apk add --no-cache upx 
COPY . .
RUN go build -ldflags="-s -w" -o /nitter-rss-proxy cmd/main.go && \
    upx --lzma /nitter-rss-proxy

FROM alpine
LABEL maintainer="github.com/dezhishen/nitter-rss-proxy"
EXPOSE 8080/tcp
WORKDIR /data
VOLUME /data
COPY --from=builder /nitter-rss-proxy /nitter-rss-proxy
RUN chmod +x /nitter-rss-proxy
ENTRYPOINT [ "/nitter-rss-proxy" ]
CMD ["-addr", "0.0.0.0:8080"]