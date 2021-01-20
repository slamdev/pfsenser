FROM alpine:3.13.0

COPY golang-cli /usr/bin/

ENTRYPOINT ["golang-cli"]
