FROM alpine:3.18

COPY gateway /usr/bin/gateway

EXPOSE 8080
ENTRYPOINT ["gateway"]