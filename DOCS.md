---
date: 2019-11-23T00:00:00+00:00
name: Git Push
title: Git Push
description: Commit and push to an git repo via SSH
authors: appleboy
author: appleboy
tags: [ deploy, publish, git-push ]
repo: appleboy/drone-git-push
logo: git.svg
icon: https://raw.githubusercontent.com/appleboy/drone-git-push/master/images/logo.svg
image: appleboy/drone-git-push
containerImage: appleboy/drone-git-push
containerImageUrl: https://hub.docker.com/r/appleboy/drone-git-push
url: https://github.com/appleboy/drone-git-push
---

Use this plugin for commit and push an git repo.
You will need to supply Drone / Woodpecker with a private SSH key or use the same credentials as the cloned repo to being able to push changes.

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    branch: master
    remote: ssh://git@git.heroku.com/falling-wind-1624.git
    force: false
    commit: true
```

An example of pushing a branch back to the current repository:

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    remote_name: origin
    branch: gh-pages
    local_ref: gh-pages
```

An example of specifying the path to a repo:

```yaml
- name: push commit
  image: appleboy/drone-git-push
  settings:
    remote_name: origin
    branch: gh-pages
    local_ref: gh-pages
    path: path/to/repo
```

An example of mirroring all refs to a remote repository:

```yaml
- name: mirror push
  image: appleboy/drone-git-push
  settings:
    remote: git@github.com:foo/bar-mirror.git
    mirror: true
    ssh_key:
      from_secret: deploy_key
```

## Parameter Reference

| setting        | description
|----------------|--------------
| ssh_key        | private SSH key for the remote machine (make sure it ends with a newline)
| remote         | target remote repository (if blank, assume exists)
| remote_name    | name of the remote to use locally (default "deploy")
| branch         | target remote branch, defaults to master
| local_branch   | local branch or ref to push (default "HEAD")
| path           | path to git repo (if blank, assume current directory)
| force          | force push using the `--force` flag, defaults to false
| skip_verify    | skip verification of HTTPS certs, defaults to false
| commit         | add and commit the contents of the repo before pushing, defaults to false
| commit_message | add a custom message for commit, if it is omitted, it will be `[skip ci] Commit dirty state`
| empty_commit   | if you only want generate an empty commit, you can do it using this option
| tag            | if you want to add a tag to the commit, you can do it using this option. You must also set `followtags` to `true` if you want the tag to be pushed to the remote
| author_name    | the name to use for the author of the commit (if blank, assume push commiter name)
| author_email   | the email address to use for the author of the commit (if blank, assume push commiter name)
| followtags     | push with --follow-tags option
| rebase         | pull --rebase before pushing
| mirror         | push all refs to remote repository with --mirror, defaults to false
