package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/drone/drone-go/drone"
)

const netrcFile = `
machine %s
login %s
password %s
`

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

// Writes the netrc file.
func WriteNetrc(in *drone.Workspace) error {
	if in.Netrc == nil || len(in.Netrc.Machine) == 0 {
		return nil
	}
	out := fmt.Sprintf(
		netrcFile,
		in.Netrc.Machine,
		in.Netrc.Login,
		in.Netrc.Password,
	)
	home := "/root"
	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}
	path := filepath.Join(home, ".netrc")
	return ioutil.WriteFile(path, []byte(out), 0600)
}
