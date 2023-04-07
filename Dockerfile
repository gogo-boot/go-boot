FROM scratch
COPY go-web-template /
ENTRYPOINT ["/app/go-web-template"]
