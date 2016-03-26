package main

type Params struct {
	Remote     string `json:"remote"`
	Branch     string `json:"branch"`
	Force      bool   `json:"force"`
	SkipVerify bool   `json:"skip_verify"`
	Commit     bool   `json:"commit"`
}
