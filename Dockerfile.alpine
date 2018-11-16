FROM alpine:3.8

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Drone Git Push" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache ca-certificates git openssh curl perl && \
  rm -rf /var/cache/apk/*

ADD release/linux/amd64/drone-git-push /bin/

ENTRYPOINT ["/bin/drone-git-push"]
