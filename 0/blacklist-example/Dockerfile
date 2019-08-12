FROM golang:latest as Builder

WORKDIR /go/src/github.com/AlexsJones/blacklist-example
COPY . .

RUN set -x && \
    export GOPATH="/go" && export PATH=$PATH:$GOPATH/bin && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh && \
    dep ensure -v && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X 'main.version=$(cat VERSION)' -X 'main.revision=$(git rev-parse --short HEAD)' -X 'main.buildtime=$(date -u +%Y-%m-%d.%H:%M:%S)'" -o /blacklist-example && \
    rm -rf ${GOPATH}

FROM alpine:3.8

COPY --from=Builder /blacklist-example /usr/bin/blacklist-example

RUN apk --no-cache add dumb-init
