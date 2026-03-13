FROM scratch
COPY web /web
ENTRYPOINT ["/web"]
