# drone-git-push

[English](README.md) | [简体中文](README.zh-cn.md)

[![GoDoc](https://godoc.org/github.com/appleboy/drone-git-push?status.svg)](https://godoc.org/github.com/appleboy/drone-git-push)
[![Lint and Testing](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml)
[![codecov](https://codecov.io/gh/appleboy/drone-git-push/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-git-push)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-git-push)](https://goreportcard.com/report/github.com/appleboy/drone-git-push)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-git-push.svg)](https://hub.docker.com/r/appleboy/drone-git-push/)

[Drone](https://www.drone.io/) / [Woodpecker](https://woodpecker-ci.org/) 插件，用於將更改推送到遠端 `git` 儲存庫。
有關使用資訊和可用選項的列表，請參閱[文件](DOCS.md)。

## 建置

使用以下命令建置二進位檔：

```sh
go build
go test
```

## Docker

使用以下命令建置 Docker 映像檔：

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-push
docker build --rm -t appleboy/drone-git-push .
```

## 使用方式

從工作目錄執行：

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

鏡像所有 refs 到遠端儲存庫：

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
