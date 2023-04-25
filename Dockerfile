FROM gcr.io/distroless/static-debian11
COPY config/config.yml config/config.yml
COPY template/* template/
COPY go-boot /
ENTRYPOINT ["/go-boot"]