package terraformcli

import (
	"bytes"
	"fmt"
	"os/exec"
)

// TODO: Improve interface documentation
// (Give each argument a name)
// (Document each method)

type Stater interface {
	List(ListOptions) ([]string, error)
	Move(string, string, MoveOptions) error
	Pull(PullOptions) ([]byte, error)
	Push(string, PushOptions) error
}

type RunnerStater struct {
	runner Runner
}

type ListOptions struct {
	// TODO: Populate
}

func (o *RunnerStater) List(options ListOptions) ([]string, error) {
	var stdout bytes.Buffer
	cmd := exec.Cmd{
		Args: []string{
			"terraform",
			"state",
			"list",
		},
		Stdout: &stdout,
	}

	if err := o.runner.Run(&cmd); err != nil {
		return nil, fmt.Errorf("error running command: %w", err)
	}

	data := stdout.Bytes()
	lines := bytes.Split(bytes.TrimSpace(data), []byte("\n"))

	resources := make([]string, len(lines))
	for i, line := range lines {
		resources[i] = string(line)
	}

	return resources, nil
}

type MoveOptions struct {
	State    string
	StateOut string

	// TODO: Add more fields
}

func (o *RunnerStater) Move(src string, dst string, options MoveOptions) error {
	args := []string{
		"terraform",
		"state",
		"mv",
	}
	if options.State != "" {
		args = append(args, fmt.Sprintf("-state=%s", options.State))
	}
	if options.StateOut != "" {
		args = append(args, fmt.Sprintf("-state-out=%s", options.StateOut))
	}
	args = append(args, src, dst)

	cmd := exec.Cmd{
		Args: args,
	}

	if err := o.runner.Run(&cmd); err != nil {
		return fmt.Errorf("error running command: %w", err)
	}

	return nil
}

type PullOptions struct {
	// TODO: Populate
}

func (o *RunnerStater) Pull(options PullOptions) ([]byte, error) {
	var stdout bytes.Buffer
	cmd := exec.Cmd{
		Args: []string{
			"terraform",
			"state",
			"pull",
		},
		Stdout: &stdout,
	}

	if err := o.runner.Run(&cmd); err != nil {
		return nil, fmt.Errorf("error running command: %w", err)
	}

	return stdout.Bytes(), nil
}

type PushOptions struct {
	// TODO: Populate
}

func (o *RunnerStater) Push(path string, options PushOptions) error {
	cmd := exec.Cmd{
		Args: []string{
			"terraform",
			"state",
			"push",
			path,
		},
	}

	if err := o.runner.Run(&cmd); err != nil {
		return fmt.Errorf("error running command: %w", err)
	}

	return nil
}

func NewRunnerStater(runner Runner) *RunnerStater {
	return &RunnerStater{runner}
}
