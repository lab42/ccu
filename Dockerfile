FROM alpine:3.20

RUN addgroup --gid 10000 -S appgroup && \
    adduser --uid 10000 -S appuser -G appgroup

COPY ccu /ccu

USER appuser

ENTRYPOINT ["/ccu"]
