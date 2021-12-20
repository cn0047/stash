FROM cn007b/go:1.16-alpine AS build

WORKDIR  /go/src/github.com/cn007b/servicetemplate
COPY . .
RUN make build

FROM alpine:3.15
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/cn007b/servicetemplate/app /usr/local/bin/

CMD ["/usr/local/bin/app"]
