FROM cn007b/alpine:latest

MAINTAINER V. K. <cn007b@gmail.com>

RUN apk add consul

COPY app /app/app

EXPOSE 8080
CMD [ "/app/app" ]
