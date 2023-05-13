FROM alpine:3.17

ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Git Push" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

LABEL org.opencontainers.image.source=https://github.com/appleboy/drone-git-push
LABEL org.opencontainers.image.description="Drone Git Push"
LABEL org.opencontainers.image.licenses=MIT

RUN apk add --no-cache ca-certificates=20230506-r0 git=2.38.5-r0 openssh=9.1_p1-r2 curl=8.0.1-r0 perl=5.36.0-r1 && \
  rm -rf /var/cache/apk/*

COPY release/${TARGETOS}/${TARGETARCH}/drone-git-push /bin/

ENTRYPOINT ["/bin/drone-git-push"]
