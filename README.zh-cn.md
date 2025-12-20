# drone-git-push

[English](README.md) | [繁體中文](README.zh-tw.md)

[![GoDoc](https://godoc.org/github.com/appleboy/drone-git-push?status.svg)](https://godoc.org/github.com/appleboy/drone-git-push)
[![Lint and Testing](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml)
[![codecov](https://codecov.io/gh/appleboy/drone-git-push/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-git-push)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-git-push)](https://goreportcard.com/report/github.com/appleboy/drone-git-push)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-git-push.svg)](https://hub.docker.com/r/appleboy/drone-git-push/)

[Drone](https://www.drone.io/) / [Woodpecker](https://woodpecker-ci.org/) 插件，用于将更改推送到远程 `git` 仓库。
有关使用信息和可用选项的列表，请查看[文档](DOCS.md)。

## 构建

使用以下命令构建二进制文件：

```sh
go build
go test
```

## Docker

使用以下命令构建 Docker 镜像：

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-push
docker build --rm -t appleboy/drone-git-push .
```

## 使用方式

从工作目录执行：

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

镜像所有 refs 到远程仓库：

```sh
docker run --rm \
  -e DRONE_COMMIT_AUTHOR=Octocat \
  -e DRONE_COMMIT_AUTHOR_EMAIL=octocat@github.com \
  -e PLUGIN_SSH_KEY="$(cat "${HOME}/.ssh/id_rsa")" \
  -e PLUGIN_REMOTE=git@github.com:foo/bar.git \
  -e PLUGIN_MIRROR=true \
  -v "$(pwd):$(pwd)" \
  -w "$(pwd)" \
  appleboy/drone-git-push
```
