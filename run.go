package terraformcli

import "os/exec"

type Runner interface {
	Run(*exec.Cmd) error
}

type mockRunner struct {
	callsRun []*mockRunnerCallRun

	writeStdout []byte
	returnErr   error
}

type mockRunnerCallRun struct {
	// The *exec.Cmd passed to the Run() function.
	//
	// Note that only the exec.Cmd.Args field is populated. All other fields are zeroed.
	cmd *exec.Cmd
}

func (o *mockRunner) Run(cmd *exec.Cmd) error {
	call := mockRunnerCallRun{
		cmd: &exec.Cmd{
			Args: cmd.Args,
		},
	}
	o.callsRun = append(o.callsRun, &call)

	if cmd.Stdout != nil {
		cmd.Stdout.Write(o.writeStdout)
	}

	return o.returnErr
}
