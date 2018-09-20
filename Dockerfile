FROM golang:1.11-alpine

ADD . /go/src/github.com/banzaicloud/satellite
WORKDIR /go/src/github.com/banzaicloud/satellite
RUN apk update && apk add ca-certificates make git curl

RUN make vendor
RUN go build -o /tmp/satellite main.go

FROM alpine:3.7

COPY --from=0 /tmp/satellite /usr/local/bin/satellite
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN adduser -D satellite

USER satellite

ENTRYPOINT ["/usr/local/bin/satellite"]
