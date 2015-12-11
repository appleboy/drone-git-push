# Docker image for the Drone Git Push plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-git-push
#     make deps build
#     docker build --rm=true -t plugins/drone-git-push .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    git \
    openssh \
    curl \
    perl && \
  rm -rf /var/cache/apk/*

ADD drone-git-push /bin/
ENTRYPOINT ["/bin/drone-git-push"]
