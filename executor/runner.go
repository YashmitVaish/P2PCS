package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Runner struct {
	code    string
	WorkDir string
}

func GetRunner(code string) (*Runner, error) {
	dir, err := os.MkdirTemp("", "p2pcs-exec-*")
	if err != nil {
		return nil, err
	}

	return &Runner{
		code:    code,
		WorkDir: dir,
	}, nil
}

func (r *Runner) Run(ctx context.Context) (int, error) {
	defer os.RemoveAll(r.WorkDir)

	var cmd *exec.Cmd

	filename := "main.py"
	err := os.WriteFile(filepath.Join(r.WorkDir, filename), []byte(r.code), 0644)
	if err != nil {
		fmt.Println("rn==36")
		return -1, err
	}
	cmd = exec.CommandContext(ctx, "python", "-u", filename)

	cmd.Dir = r.WorkDir

	stdout, stderr, err := AttachStream(cmd)

	if err != nil {
		fmt.Println("rn==46")
		return -1, err
	}

	err = cmd.Start()
	if err != nil {
		return -1, err
	}

	go StreamOutput("STDOUT", stdout)
	go StreamOutput("STDERR", stderr)

	err = cmd.Wait()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("rn==56")
		return -1, fmt.Errorf("execution timed out")
	}

	if err != nil {
		fmt.Println("rn==61")
		if exitErr, ok := err.(*exec.ExitError); ok {
			fmt.Println("rn==63")
			return exitErr.ExitCode(), nil
		}
		return -1, err
	}

	return 0, nil
}
