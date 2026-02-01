package main

import (
	"context"
	"fmt"
)

type ExecResult struct {
	ExitCode int
}

func RunCode(ctx context.Context, code string) (*ExecResult, error) {
	runner, err := GetRunner(code)
	if err != nil {
		fmt.Println("ex==15")
		return nil, err
	}

	exitCode, err := runner.Run(ctx)
	if err != nil {
		fmt.Println("ex==20")
		return nil, err
	}

	return &ExecResult{
		ExitCode: exitCode,
	}, nil

}
