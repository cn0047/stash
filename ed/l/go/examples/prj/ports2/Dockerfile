FROM golang:1.19-alpine3.16 AS build
RUN apk add git make

WORKDIR  /go/src/github.com/ptt
COPY . .
RUN make build

FROM alpine:3.16
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/ptt/ptt /usr/local/bin/

CMD ["/usr/local/bin/ptt"]
