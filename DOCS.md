> This plugin has not been fully tested. Proceed with caution.

Use this plugin to `git push` to a repository at the end of a successful build. Use the following parameters to configure this plugin:

* `remote` - push to this remote repository
* `branch` - push to this remote branch
* `force` - force push using the `--force` flag

Example configuration in the `.drone.yml` file:

```yaml
deploy:
  git_push:
    branch: master
    remote: git@git.heroku.com:falling-wind-1624.git
    force: false
```
