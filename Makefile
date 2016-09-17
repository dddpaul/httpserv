.PHONY: all build release

IMAGE=dddpaul/httpserv
VERSION=$(shell cat VERSION)

all: build

build:
	@go test
	@mkdir -p root/bin
	@CGO_ENABLED=0 go build -o root/bin/httpserv
	@docker build --tag=${IMAGE} .

debug:
	@docker run -it --entrypoint=sh ${IMAGE}

release: build
	@docker build --tag=${IMAGE}:${VERSION} .

deploy: release
	@docker push ${IMAGE}
	@docker push ${IMAGE}:${VERSION}
