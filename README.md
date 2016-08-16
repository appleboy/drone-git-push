# drone-git-push

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-git-push/status.svg)](http://beta.drone.io/drone-plugins/drone-git-push)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-git-push?status.svg)](http://godoc.org/github.com/drone-plugins/drone-git-push)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-git-push)](https://goreportcard.com/report/github.com/drone-plugins/drone-git-push)
[![Join the chat at https://gitter.im/drone/drone](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone)

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
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t plugins/git-push .
```

Please note incorrectly building the image for the correct x64 linux and with
GCO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/drone-git-push' not found or does not exist..
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
  plugins/git-push
```
