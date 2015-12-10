> This plugin has not been fully tested. Proceed with caution.

Use this plugin for deplying an application via `git push`. You can override
the default configuration with the following parameters:

* `remote` - Target remote repository
* `branch` - Target remote branch, defaults to master
* `force` - Force push using the `--force` flag, defaults to false
* `skip_verify` - Skip verification of HTTPS certs, defaults to false

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  git_push:
    branch: master
    remote: git@git.heroku.com:falling-wind-1624.git
    force: false
```
