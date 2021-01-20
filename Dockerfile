FROM alpine:3.13.0

COPY pfsenser /usr/bin/

ENTRYPOINT ["pfsenser"]
