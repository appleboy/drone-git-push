# drone-git-push

[English](README.md) | [简体中文](README.zh-cn.md)

[![GoDoc](https://godoc.org/github.com/appleboy/drone-git-push?status.svg)](https://godoc.org/github.com/appleboy/drone-git-push)
[![Lint and Testing](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml)
[![codecov](https://codecov.io/gh/appleboy/drone-git-push/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-git-push)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-git-push)](https://goreportcard.com/report/github.com/appleboy/drone-git-push)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-git-push.svg)](https://hub.docker.com/r/appleboy/drone-git-push/)

支援 [Drone](https://www.drone.io/)、[Woodpecker](https://woodpecker-ci.org/)、[GitHub Actions](https://github.com/features/actions) 和 [Gitea Actions](https://docs.gitea.com/usage/actions/overview) 的 CI/CD 插件，用於將更改推送到遠端 Git 儲存庫。

## 目錄

- [drone-git-push](#drone-git-push)
  - [目錄](#目錄)
  - [功能特點](#功能特點)
  - [支援平台](#支援平台)
  - [使用方式](#使用方式)
    - [Drone / Woodpecker](#drone--woodpecker)
    - [GitHub Actions](#github-actions)
    - [Gitea Actions](#gitea-actions)
  - [參數說明](#參數說明)
  - [認證方式](#認證方式)
    - [SSH 金鑰](#ssh-金鑰)
    - [HTTPS 帳號密碼](#https-帳號密碼)
  - [從原始碼建置](#從原始碼建置)
  - [使用 Docker 執行](#使用-docker-執行)
  - [授權條款](#授權條款)

## 功能特點

- 透過 SSH 或 HTTPS 推送提交到遠端儲存庫
- 鏡像所有 refs 到遠端儲存庫
- 推送前自動提交變更
- 支援標籤與 follow-tags 選項
- 推送前執行 rebase
- 支援強制推送
- 自訂提交訊息
- 支援空提交
- 支援 Git LFS

## 支援平台

| CI 平台        | 狀態     |
| -------------- | -------- |
| Drone          | 完整支援 |
| Woodpecker     | 完整支援 |
| GitHub Actions | 完整支援 |
| Gitea Actions  | 完整支援 |

## 使用方式

### Drone / Woodpecker

基本推送到遠端分支：

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    branch: master
    remote: git@github.com:foo/bar.git
    ssh_key:
      from_secret: deploy_key
```

推送並提交變更：

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

推送到當前儲存庫：

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    remote_name: origin
    branch: gh-pages
    local_ref: gh-pages
```

鏡像所有 refs 到遠端儲存庫：

```yaml
- name: mirror push
  image: appleboy/drone-git-push
  settings:
    remote: git@github.com:foo/bar-mirror.git
    mirror: true
    ssh_key:
      from_secret: deploy_key
```

推送並加上標籤：

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

## 參數說明

| 參數             | 說明                           | 預設值                         |
| ---------------- | ------------------------------ | ------------------------------ |
| `ssh_key`        | 遠端主機的私有 SSH 金鑰        | -                              |
| `remote`         | 目標遠端儲存庫 URL             | -                              |
| `remote_name`    | 本地使用的遠端名稱             | `deploy`                       |
| `branch`         | 目標遠端分支                   | `master`                       |
| `local_branch`   | 要推送的本地分支或 ref         | `HEAD`                         |
| `path`           | Git 儲存庫路徑                 | 當前目錄                       |
| `force`          | 使用 `--force` 強制推送        | `false`                        |
| `skip_verify`    | 跳過 HTTPS 憑證驗證            | `false`                        |
| `commit`         | 推送前新增並提交內容           | `false`                        |
| `commit_message` | 自訂提交訊息                   | `[skip ci] Commit dirty state` |
| `empty_commit`   | 建立空提交                     | `false`                        |
| `no_verify`      | 略過 pre-commit 和 commit-msg  | `false`                        |
| `tag`            | 要加到提交的標籤               | -                              |
| `followtags`     | 使用 `--follow-tags` 選項推送  | `false`                        |
| `rebase`         | 推送前執行 `--rebase`          | `false`                        |
| `mirror`         | 使用 `--mirror` 推送所有 refs  | `false`                        |
| `author_name`    | 提交的作者名稱                 | CI 提交作者                    |
| `author_email`   | 提交的作者電子郵件             | CI 提交作者電子郵件            |

## 認證方式

### SSH 金鑰

提供私有 SSH 金鑰進行認證：

```yaml
settings:
  ssh_key:
    from_secret: deploy_key
```

### HTTPS 帳號密碼

使用 netrc 憑證進行 HTTPS 認證：

```yaml
settings:
  username:
    from_secret: git_username
  password:
    from_secret: git_password
```

## 從原始碼建置

建置二進位檔：

```sh
go build
go test
```

建置 Docker 映像檔：

```sh
# 建置 Linux amd64 版本
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-push

# 建置 Docker 映像檔
docker build --rm -t appleboy/drone-git-push -f docker/Dockerfile .
```

## 使用 Docker 執行

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

## 授權條款

MIT 授權條款 - 詳見 [LICENSE](LICENSE) 檔案。
