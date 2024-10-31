FROM alpine:3.20 as setup

RUN addgroup --gid 10000 -S appgroup && \
    adduser --uid 10000 -S appuser -G appgroup

COPY ccu /ccu

USER appuser

ENTRYPOINT ["/ccu"]
