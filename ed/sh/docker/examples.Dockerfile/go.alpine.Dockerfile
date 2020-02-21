FROM alpine:latest

MAINTAINER Vladimir Kovpak <cn007b@gmail.com>

COPY xgoapp .
CMD [ "./xgoapp" ]
EXPOSE 8080
ENTRYPOINT [ "./xgoapp" ]
