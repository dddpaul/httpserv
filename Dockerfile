FROM golang:1.13.5 as builder
WORKDIR /go/src/github.com/dddpaul/httpserv
ADD . ./
RUN make build-alpine

FROM alpine:latest
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/* && \
    update-ca-certificates && \
    addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /go/src/github.com/dddpaul/httpserv/bin/httpserv .
EXPOSE 8080

CMD ["./httpserv", "-port", ":8080"]
