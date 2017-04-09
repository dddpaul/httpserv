FROM alpine:3.3
MAINTAINER Pavel Derendyaev <dddpaul@gmail.com>

ADD root /

ENTRYPOINT ["/bin/httpserv"]
CMD ["-port", ":8080"]
EXPOSE 8080
