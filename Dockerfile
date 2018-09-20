FROM golang:1.11-alpine

ADD . /go/src/github.com/banzaicloud/noaa
WORKDIR /go/src/github.com/banzaicloud/noaa
RUN apk update && apk add ca-certificates make git curl

RUN make vendor
RUN go build -o /tmp/noaa main.go

FROM alpine:3.7

COPY --from=0 /tmp/noaa /usr/local/bin/noaa
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN adduser -D noaa

USER noaa

ENTRYPOINT ["/usr/local/bin/noaa"]
