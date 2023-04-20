FROM gcr.io/distroless/static-debian11
COPY config/config.yml config/config.yml
COPY go-web-template /
ENTRYPOINT ["/go-web-template"]