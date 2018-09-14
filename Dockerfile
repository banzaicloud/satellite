FROM golang:1.11-alpine

ADD . /go/src/github.com/banzaicloud/noaa
WORKDIR /go/src/github.com/banzaicloud/noaa
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /tmp/noaa main.go

FROM alpine:3.6

COPY --from=0 /tmp/noaa /usr/local/bin/noaa
RUN apk update && apk add ca-certificates
RUN adduser -D noaa

USER noaa

ENTRYPOINT ["/usr/local/bin/noaa"]
