package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	code := `
import time
print("hello from executor")
for i in range(60):
	print(i)
	time.sleep(1)
	`
	parent := context.Background()

	ctx, cancel := context.WithTimeout(parent, 60*time.Second)
	defer cancel()

	result, err := RunCode(ctx, code)
	if err != nil {
		fmt.Println("execution error:", err)
		return
	}

	fmt.Println("Programme Finished with code: ", result.ExitCode)
}
