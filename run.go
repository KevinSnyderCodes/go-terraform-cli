package terraformcli

import "os/exec"

type Runner interface {
	Run(*exec.Cmd) error
}

type mockRunner struct {
	writeStdout []byte

	returnErr error
}

func (o *mockRunner) Run(cmd *exec.Cmd) error {
	if cmd.Stdout != nil {
		cmd.Stdout.Write(o.writeStdout)
	}

	return o.returnErr
}
