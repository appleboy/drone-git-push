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

COPY release/${TARGETOS}/${TARGETARCH}/drone-git-push /bin/

ENTRYPOINT ["/bin/drone-git-push"]
