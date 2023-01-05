# drone-git-push

[![GoDoc](https://godoc.org/github.com/appleboy/drone-git-push?status.svg)](https://godoc.org/github.com/appleboy/drone-git-push)
[![Lint and Testing](https://github.com/JonasBernard/drone-git-push/actions/workflows/lint.yml/badge.svg)](https://github.com/JonasBernard/drone-git-push/actions/workflows/lint.yml)
[![codecov](https://codecov.io/gh/JonasBernard/drone-git-push/branch/master/graph/badge.svg)](https://codecov.io/gh/JonasBernard/drone-git-push)
[![Go Report Card](https://goreportcard.com/badge/github.com/JonasBernard/drone-git-push)](https://goreportcard.com/report/github.com/JonasBernard/drone-git-push)
[![Docker Pulls](https://img.shields.io/docker/pulls/jonasbernard/drone-git-push.svg)](https://hub.docker.com/r/jonasbernard/drone-git-push/)

[Drone](https://www.drone.io/) / [Woodpecker](https://woodpecker-ci.org/) plugin to push changes to a remote `git` repository.
For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Build

Build the binary with the following commands:

```sh
go build
go test
```

## Docker

Build the docker image with the following commands:

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-push
docker build --rm -t appleboy/drone-git-push .
```

## Usage

Execute from the working directory:

```sh
docker run --rm \
  -e DRONE_COMMIT_AUTHOR=Octocat \
  -e DRONE_COMMIT_AUTHOR_EMAIL=octocat@github.com \
  -e PLUGIN_SSH_KEY="$(cat "${HOME}/.ssh/id_rsa")" \
  -e PLUGIN_BRANCH=master \
  -e PLUGIN_REMOTE=git@github.com:foo/bar.git \
  -e PLUGIN_FORCE=false \
  -v "$(pwd):$(pwd)" \
  -w "$(pwd)" \
  appleboy/drone-git-push
```
