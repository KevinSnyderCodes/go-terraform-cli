package terraformcli

import (
	"fmt"
	"os/exec"
)

type CLI struct {
	path string
	dir  string

	State Stater
}

func (o *CLI) Run(cmd *exec.Cmd) error {
	cmd.Path = o.path
	cmd.Dir = o.dir

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running command: %w", err)
	}

	return nil
}

func NewCLI(dir string) (*CLI, error) {
	path, err := exec.LookPath("terraform")
	if err != nil {
		return nil, fmt.Errorf("error finding executable in path: %w", err)
	}

	cli := CLI{path: path, dir: dir}
	cli.State = NewRunnerStater(&cli)

	return &cli, nil
}
