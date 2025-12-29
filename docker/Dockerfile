FROM alpine:3.23

ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Git Push" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

LABEL org.opencontainers.image.source=https://github.com/appleboy/drone-git-push
LABEL org.opencontainers.image.description="Drone Git Push"
LABEL org.opencontainers.image.licenses=MIT

RUN apk add --no-cache ca-certificates git git-lfs openssh curl perl && \
  rm -rf /var/cache/apk/*

RUN adduser -D -u 1000 appuser

# Configure Git to trust all directories to avoid "fatal: not in a git directory" errors
# This is necessary when running as non-root user in CI/CD environments
RUN git config --system --add safe.directory '*'

COPY release/${TARGETOS}/${TARGETARCH}/drone-git-push /bin/

USER appuser

ENTRYPOINT ["/bin/drone-git-push"]
