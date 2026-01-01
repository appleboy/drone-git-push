# drone-git-push

[繁體中文](README.zh-tw.md) | [简体中文](README.zh-cn.md)

[![GoDoc](https://godoc.org/github.com/appleboy/drone-git-push?status.svg)](https://godoc.org/github.com/appleboy/drone-git-push)
[![Lint and Testing](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/testing.yml)
[![Trivy Security Scan](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml/badge.svg)](https://github.com/appleboy/drone-git-push/actions/workflows/trivy.yml)
[![codecov](https://codecov.io/gh/appleboy/drone-git-push/branch/master/graph/badge.svg)](https://codecov.io/gh/appleboy/drone-git-push)
[![Go Report Card](https://goreportcard.com/badge/github.com/appleboy/drone-git-push)](https://goreportcard.com/report/github.com/appleboy/drone-git-push)
[![Docker Pulls](https://img.shields.io/docker/pulls/appleboy/drone-git-push.svg)](https://hub.docker.com/r/appleboy/drone-git-push/)

A CI/CD plugin for [Drone](https://www.drone.io/), [Woodpecker](https://woodpecker-ci.org/), [Crow CI](https://crowci.dev/), [GitHub Actions](https://github.com/features/actions), and [Gitea Actions](https://docs.gitea.com/usage/actions/overview) to push changes to a remote Git repository.

## Table of Contents

- [drone-git-push](#drone-git-push)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Supported Platforms](#supported-platforms)
  - [Usage](#usage)
    - [Drone / Woodpecker](#drone--woodpecker)
    - [GitHub Actions](#github-actions)
    - [Gitea Actions](#gitea-actions)
  - [Parameter Reference](#parameter-reference)
  - [Authentication](#authentication)
    - [SSH Key](#ssh-key)
    - [HTTPS with Username/Password](#https-with-usernamepassword)
  - [Build from Source](#build-from-source)
  - [Run with Docker](#run-with-docker)
  - [License](#license)

## Features

- Push commits to remote repositories via SSH or HTTPS
- Mirror all refs to a remote repository
- Auto-commit dirty changes before pushing
- Tag support with follow-tags option
- Rebase before push
- Force push support
- Custom commit messages
- Empty commit support
- Git LFS support

## Supported Platforms

| CI Platform    | Status          |
| -------------- | --------------- |
| Drone          | Fully supported |
| Woodpecker     | Fully supported |
| GitHub Actions | Fully supported |
| Gitea Actions  | Fully supported |

## Usage

### Drone / Woodpecker

Basic push to a remote branch:

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    branch: master
    remote: git@github.com:foo/bar.git
    ssh_key:
      from_secret: deploy_key
```

Push with commit changes:

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

Push to the current repository:

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    remote_name: origin
    branch: gh-pages
    local_ref: gh-pages
```

Mirror all refs to a remote repository:

```yaml
- name: mirror push
  image: appleboy/drone-git-push
  settings:
    remote: git@github.com:foo/bar-mirror.git
    mirror: true
    ssh_key:
      from_secret: deploy_key
```

Push with tagging:

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

## Parameter Reference

| Parameter        | Description                                | Default                        |
| ---------------- | ------------------------------------------ | ------------------------------ |
| `ssh_key`        | Private SSH key for the remote machine     | -                              |
| `remote`         | Target remote repository URL               | -                              |
| `remote_name`    | Name of the remote to use locally          | `deploy`                       |
| `branch`         | Target remote branch                       | `master`                       |
| `local_branch`   | Local branch or ref to push                | `HEAD`                         |
| `path`           | Path to git repository                     | Current directory              |
| `force`          | Force push using `--force` flag            | `false`                        |
| `skip_verify`    | Skip verification of HTTPS certs           | `false`                        |
| `commit`         | Add and commit the contents before pushing | `false`                        |
| `commit_message` | Custom commit message                      | `[skip ci] Commit dirty state` |
| `empty_commit`   | Create an empty commit                     | `false`                        |
| `no_verify`      | Bypass pre-commit and commit-msg hooks     | `false`                        |
| `tag`            | Tag to add to the commit                   | -                              |
| `followtags`     | Push with `--follow-tags` option           | `false`                        |
| `rebase`         | Pull `--rebase` before pushing             | `false`                        |
| `mirror`         | Push all refs with `--mirror`              | `false`                        |
| `author_name`    | Author name for the commit                 | CI commit author               |
| `author_email`   | Author email for the commit                | CI commit author email         |

## Authentication

### SSH Key

Provide a private SSH key for authentication:

```yaml
settings:
  ssh_key:
    from_secret: deploy_key
```

### HTTPS with Username/Password

Use netrc credentials for HTTPS authentication:

```yaml
settings:
  username:
    from_secret: git_username
  password:
    from_secret: git_password
```

## Build from Source

Build the binary:

```sh
go build
go test
```

Build Docker image:

```sh
# Build for Linux amd64
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-git-push

# Build Docker image
docker build --rm -t appleboy/drone-git-push -f docker/Dockerfile .
```

## Run with Docker

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

## License

MIT License - see the [LICENSE](LICENSE) file for details.
