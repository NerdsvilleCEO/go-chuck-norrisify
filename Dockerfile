FROM golang:1.9

WORKDIR /go/src/github.com/NerdsvilleCEO/go-chuck-norrisify
COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go get ./...

COPY docker-entrypoint.sh /usr/bin
ENTRYPOINT ["docker-entrypoint.sh"]
