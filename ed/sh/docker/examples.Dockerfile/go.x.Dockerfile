FROM cn007b/go

MAINTAINER V. Kovpak <cn007b@gmail.com>

COPY xgoapp .
CMD [ "/bin/bash", "-c", "./xgoapp" ]
EXPOSE 8080
ENTRYPOINT [ "/bin/bash", "-c", "./xgoapp" ]
