package main

type Params struct {
	Remote      string `json:"remote"`
	RemoteName  string `json:"remote_name"`
	Branch      string `json:"branch"`
	LocalBranch string `json:"local_branch"`
	Force       bool   `json:"force"`
	SkipVerify  bool   `json:"skip_verify"`
	Commit      bool   `json:"commit"`
}
