FROM gcr.io/distroless/static-debian11
COPY platform/config/config.yml config/config.yml
COPY web/template/* template/
COPY go-boot /
ENTRYPOINT ["/go-boot"]