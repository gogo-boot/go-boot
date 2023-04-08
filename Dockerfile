FROM scratch
COPY go-web-template /
ENTRYPOINT ["/go-web-template"]
