FROM golang:1.13.4 as builder
WORKDIR /go/src/github.com/dddpaul/httpserv
ADD . ./
RUN make build-alpine

FROM alpine:latest
RUN apk add --update ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/* && \
    update-ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/dddpaul/httpserv/bin/httpserv .

ENTRYPOINT ["./httpserv"]
CMD ["-port", ":8080"]
EXPOSE 8080
