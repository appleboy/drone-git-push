package repo

import (
  "path/filepath"
  "os/user"
  "io/ioutil"
  "os"

  "github.com/drone/drone-go/drone"
)

func WriteKey(workspace *drone.Workspace) error {
  if workspace.Keys == nil || len(workspace.Keys.Private) == 0 {
    return nil
  }

  home := "/root"

  if currentUser, err := user.Current(); err == nil {
    home = currentUser.HomeDir
  }

  sshpath := filepath.Join(
    home,
    ".ssh")

  if err := os.MkdirAll(sshpath, 0700); err != nil {
    return err
  }

  confpath := filepath.Join(
    sshpath,
    "config")

  privpath := filepath.Join(
    sshpath,
    "id_rsa")

  ioutil.WriteFile(
    confpath,
    []byte("StrictHostKeyChecking no\n"),
    0700)

  return ioutil.WriteFile(
    privpath,
    []byte(workspace.Keys.Private),
    0600)
}
