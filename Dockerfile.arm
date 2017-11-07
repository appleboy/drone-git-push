FROM plugins/base:arm

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/appleboy/drone-git-push.git"
LABEL org.label-schema.name="drone git push plugin"
LABEL org.label-schema.vendor="Bo-Yi Wu"
LABEL org.label-schema.schema-version="1.0"
LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>"

RUN apk add --no-cache ca-certificates git openssh curl perl && \
  rm -rf /var/cache/apk/*

ADD release/linux/arm/drone-git-push /bin/
ENTRYPOINT ["/bin/drone-git-push"]
