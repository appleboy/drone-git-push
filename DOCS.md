Use this plugin for deplying an application via `git push`. You will need to
supply Drone with a private SSH key or use the same credentials as the cloned
repo to being able to push changes.

## Config

The following parameters are used to configure the plugin:

* **key** - private SSH key for the remote machine
* **remote** - target remote repository (if blank, assume exists)
* **remote_name** - name of the remote to use locally (default "deploy")
* **branch** - target remote branch, defaults to master
* **local_branch** - local branch or ref to push (default "HEAD")
* **force** - force push using the `--force` flag, defaults to false
* **skip_verify** - skip verification of HTTPS certs, defaults to false
* **commit** - add and commit the contents of the repo before pushing, defaults to false

The following secret values can be set to configure the plugin.

* **GIT_PUSH_SSH_KEY** - corresponds to **key**

It is highly recommended to put the **GIT_PUSH_SSH_KEY** into a secret so it is
not exposed to users. This can be done using the drone-cli.

```bash
drone secret add --image=plugins/git-push \
    octocat/hello-world GIT_PUSH_SSH_KEY @path/to/.ssh/id_rsa
```

Then sign the YAML file after all secrets are added.

```bash
drone sign octocat/hello-world
```

See [secrets](http://readme.drone.io/0.5/usage/secrets/) for additional
information on secrets

## Examples

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:
  git_push:
    image: plugins/git-push
    branch: master
    remote: git@git.heroku.com:falling-wind-1624.git
    force: false
    commit: true
```

An example of pushing a branch back to the current repository:

```yaml
pipeline:
  git_push:
    image: plugins/git-push
    remote_name: origin
    branch: gh-pages
    local_ref: gh-pages
```
