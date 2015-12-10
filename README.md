# drone-git-push

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-git-push/status.svg)](http://beta.drone.io/drone-plugins/drone-git-push)
[![](https://badge.imagelayers.io/plugins/drone-git-push:latest.svg)](https://imagelayers.io/?images=plugins/drone-git-push:latest 'Get your own badge on imagelayers.io')

Drone plugin for deploying via Git

## Usage

```sh
./drone-git-push <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "branch": "master",
        "remote": "git@git.heroku.com:falling-wind-1624.git",
        "force: false"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```sh
make deps build
docker build --rm=true -t plugins/drone-git-push .
```

### Example

```sh
docker run -i plugins/drone-git-push <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "branch": "master",
        "remote": "git@git.heroku.com:falling-wind-1624.git",
        "force: false"
    }
}
EOF
```
