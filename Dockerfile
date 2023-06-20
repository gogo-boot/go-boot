FROM gcr.io/distroless/static-debian11
COPY platform/config/* platform/config/
COPY web/template/* web/template/
COPY go-boot /
ENTRYPOINT ["/go-boot"]