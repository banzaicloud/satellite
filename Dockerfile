FROM golang:1.11-alpine

ADD . /go/src/github.com/banzaicloud/whereami
WORKDIR /go/src/github.com/banzaicloud/whereami
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /tmp/whereami main.go

FROM alpine:3.6

COPY --from=0 /tmp/whereami /usr/local/bin/whereami
RUN apk update && apk add ca-certificates
RUN adduser -D whereami

USER whereami

ENTRYPOINT ["/usr/local/bin/whereami"]