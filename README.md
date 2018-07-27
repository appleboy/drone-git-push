# drone-git-push

[![Build Status](http://drone.wu-boy.com/api/badges/appleboy/drone-git-push/status.svg)](http://drone.wu-boy.com/appleboy/drone-git-push)
[![Go Doc](https://godoc.org/github.com/appleboy/drone-git-push?status.svg)](http://godoc.org/github.com/appleboy/drone-git-push)
[![Go Report](https://goreportcard.com/badge/github.com/appleboy/drone-git-push)](https://goreportcard.com/report/github.com/appleboy/drone-git-push)

Drone plugin to push changes to a remote `git` repository. For the usage
information and a listing of the available options please take a look at
[the docs](DOCS.md).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-push
docker build --rm -t appleboy/drone-git-push .
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
  -e DRONE_COMMIT_AUTHOR=Octocat \
  -e DRONE_COMMIT_AUTHOR_EMAIL=octocat@github.com \
  -e PLUGIN_SSH_KEY=${HOME}/.ssh/id_rsa \
  -e PLUGIN_BRANCH=master \
  -e PLUGIN_REMOTE=git@github.com:foo/bar.git \
  -e PLUGIN_FORCE=false \
  -v $(pwd)/$(pwd) \
  -w $(pwd) \
  appleboy/drone-git-push
```
