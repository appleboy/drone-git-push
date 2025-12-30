# drone-git-push

[English](README.md) | [繁體中文](README.zh-tw.md)

[![GoDoc](https://godoc.org/github.com/appleboy/drone-git-push?status.svg)](https://godoc.org/github.com/appleboy/drone-git-push)
[![Lint and Testing](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml)
[![codecov](https://codecov.io/gh/appleboy/drone-git-push/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-git-push)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-git-push)](https://goreportcard.com/report/github.com/appleboy/drone-git-push)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-git-push.svg)](https://hub.docker.com/r/appleboy/drone-git-push/)

支持 [Drone](https://www.drone.io/)、[Woodpecker](https://woodpecker-ci.org/)、[GitHub Actions](https://github.com/features/actions) 和 [Gitea Actions](https://docs.gitea.com/usage/actions/overview) 的 CI/CD 插件，用于将更改推送到远程 Git 仓库。

## 目录

- [drone-git-push](#drone-git-push)
  - [目录](#目录)
  - [功能特点](#功能特点)
  - [支持平台](#支持平台)
  - [使用方式](#使用方式)
    - [Drone / Woodpecker](#drone--woodpecker)
    - [GitHub Actions](#github-actions)
    - [Gitea Actions](#gitea-actions)
  - [参数说明](#参数说明)
  - [认证方式](#认证方式)
    - [SSH 密钥](#ssh-密钥)
    - [HTTPS 账号密码](#https-账号密码)
  - [从源码构建](#从源码构建)
  - [使用 Docker 运行](#使用-docker-运行)
  - [许可证](#许可证)

## 功能特点

- 通过 SSH 或 HTTPS 推送提交到远程仓库
- 镜像所有 refs 到远程仓库
- 推送前自动提交更改
- 支持标签与 follow-tags 选项
- 推送前执行 rebase
- 支持强制推送
- 自定义提交信息
- 支持空提交
- 支持 Git LFS

## 支持平台

| CI 平台        | 状态     |
| -------------- | -------- |
| Drone          | 完整支持 |
| Woodpecker     | 完整支持 |
| GitHub Actions | 完整支持 |
| Gitea Actions  | 完整支持 |

## 使用方式

### Drone / Woodpecker

基本推送到远程分支：

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    branch: master
    remote: git@github.com:foo/bar.git
    ssh_key:
      from_secret: deploy_key
```

推送并提交更改：

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    branch: master
    remote: git@github.com:foo/bar.git
    force: false
    commit: true
    commit_message: "[skip ci] Update generated files"
    ssh_key:
      from_secret: deploy_key
```

推送到当前仓库：

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    remote_name: origin
    branch: gh-pages
    local_ref: gh-pages
```

镜像所有 refs 到远程仓库：

```yaml
- name: mirror push
  image: appleboy/drone-git-push
  settings:
    remote: git@github.com:foo/bar-mirror.git
    mirror: true
    ssh_key:
      from_secret: deploy_key
```

推送并添加标签：

```yaml
- name: push with tag
  image: appleboy/drone-git-push
  settings:
    branch: master
    remote: git@github.com:foo/bar.git
    commit: true
    tag: v1.0.0
    followtags: true
    ssh_key:
      from_secret: deploy_key
```

### GitHub Actions

```yaml
- name: Push changes
  uses: appleboy/drone-git-push@master
  with:
    remote: git@github.com:foo/bar.git
    branch: master
    ssh_key: ${{ secrets.DEPLOY_KEY }}
```

### Gitea Actions

```yaml
- name: Push changes
  uses: appleboy/drone-git-push@master
  with:
    remote: git@gitea.com:foo/bar.git
    branch: master
    ssh_key: ${{ secrets.DEPLOY_KEY }}
```

## 参数说明

| 参数             | 说明                           | 默认值                         |
| ---------------- | ------------------------------ | ------------------------------ |
| `ssh_key`        | 远程主机的私有 SSH 密钥        | -                              |
| `remote`         | 目标远程仓库 URL               | -                              |
| `remote_name`    | 本地使用的远程名称             | `deploy`                       |
| `branch`         | 目标远程分支                   | `master`                       |
| `local_branch`   | 要推送的本地分支或 ref         | `HEAD`                         |
| `path`           | Git 仓库路径                   | 当前目录                       |
| `force`          | 使用 `--force` 强制推送        | `false`                        |
| `skip_verify`    | 跳过 HTTPS 证书验证            | `false`                        |
| `commit`         | 推送前添加并提交内容           | `false`                        |
| `commit_message` | 自定义提交信息                 | `[skip ci] Commit dirty state` |
| `empty_commit`   | 创建空提交                     | `false`                        |
| `no_verify`      | 跳过 pre-commit 和 commit-msg  | `false`                        |
| `tag`            | 要添加到提交的标签             | -                              |
| `followtags`     | 使用 `--follow-tags` 选项推送  | `false`                        |
| `rebase`         | 推送前执行 `--rebase`          | `false`                        |
| `mirror`         | 使用 `--mirror` 推送所有 refs  | `false`                        |
| `author_name`    | 提交的作者名称                 | CI 提交作者                    |
| `author_email`   | 提交的作者电子邮件             | CI 提交作者电子邮件            |

## 认证方式

### SSH 密钥

提供私有 SSH 密钥进行认证：

```yaml
settings:
  ssh_key:
    from_secret: deploy_key
```

### HTTPS 账号密码

使用 netrc 凭证进行 HTTPS 认证：

```yaml
settings:
  username:
    from_secret: git_username
  password:
    from_secret: git_password
```

## 从源码构建

构建二进制文件：

```sh
go build
go test
```

构建 Docker 镜像：

```sh
# 构建 Linux amd64 版本
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-push

# 构建 Docker 镜像
docker build --rm -t appleboy/drone-git-push -f docker/Dockerfile .
```

## 使用 Docker 运行

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

## 许可证

MIT 许可证 - 详见 [LICENSE](LICENSE) 文件。
