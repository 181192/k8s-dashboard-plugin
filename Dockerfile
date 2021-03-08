FROM alpine:3.13.2 as deps
RUN apk --update add ca-certificates
RUN mkdir tmp-dir

FROM scratch

COPY --from=deps /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=deps tmp-dir tmp

COPY dashboard /

ENTRYPOINT ["/dashboard"]