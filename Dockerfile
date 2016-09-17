FROM alpine:3.3
MAINTAINER Pavel Derendyaev <pderendyaev@smile-net.ru>

ADD root /

ENTRYPOINT ["/bin/httpserv"]
CMD ["-port", ":8080"]
EXPOSE 8080
