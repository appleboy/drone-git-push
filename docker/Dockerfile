FROM alpine:3.17

ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Git Push" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates=20220614-r4 git=2.38.5-r0 openssh=9.1_p1-r2 curl=8.0.1-r0 perl=5.36.0-r1 && \
  rm -rf /var/cache/apk/*

COPY release/${TARGETOS}/${TARGETARCH}/drone-git-push /bin/

ENTRYPOINT ["/bin/drone-git-push"]
