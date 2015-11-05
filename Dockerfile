# Docker image for Drone's git push plugin
#
#     CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-git-push .

FROM alpine:3.2
RUN apk add -U ca-certificates git openssh curl perl && rm -rf /var/cache/apk/*
ADD drone-git-push /bin/
ENTRYPOINT ["/bin/drone-git-push"]
