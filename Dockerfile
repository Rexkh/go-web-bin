FROM scratch
COPY web /web
USER 10001
ENTRYPOINT ["/web"]
