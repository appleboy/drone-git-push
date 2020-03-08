package repo

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

const netrcFile = `
machine %s
login %s
password %s
`

// WriteKey writes the private key.
func WriteKey(privateKey string) error {
	if privateKey == "" {
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
		[]byte(privateKey),
		0600)
}

// WriteNetrc writes the netrc file.
func WriteNetrc(machine, login, password string) error {
	if machine == "" {
		return nil
	}

	netrcContent := fmt.Sprintf(
		netrcFile,
		machine,
		login,
		password,
	)

	home := "/root"

	if currentUser, err := user.Current(); err == nil {
		home = currentUser.HomeDir
	}

	netpath := filepath.Join(
		home,
		".netrc")

	return ioutil.WriteFile(
		netpath,
		[]byte(netrcContent),
		0600)
}

// WriteToken authenticate with Git hosting using a token.
func WriteToken(remote string, login, password string) (string, error) {
	if remote == "" || login == "" || password == "" {
		return remote, nil
	}

	u, err := url.Parse(remote)

	if err != nil {
		return remote, err
	}

	u.User = url.UserPassword(login, password)

	return u.String(), nil
}
