FROM cn007b/alpine

MAINTAINER V. Kovpak <cn007b@gmail.com>

ENV X_URI=https://realtimelog.herokuapp.com:443/rkc8q6llprn

ENTRYPOINT ["/bin/sh", "-c", "echo \"\nPlease open: $X_URI\n\" && curl -i -XPOST $X_URI -H 'Content-Type: application/json' -d '{\"code\": 200, \"status\": \"ok\", \"version\": \"v1\"}'"]
