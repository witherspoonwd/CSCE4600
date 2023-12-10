package builtins

import (
	"fmt"
	"os/exec"
	"time"
)

func Time(args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("You must pass a command to run") // bye
	}

	startTime := time.Now()

	realArgs := args[1:] // get the actual args without the command name

	// execute command
	cmd := exec.Command(args[0], realArgs...)
	err := cmd.Run()

	// if it fails, don't kill the shell
	if err != nil {
		return fmt.Errorf("Error: ", err)
	}

	fmt.Println(time.Since(startTime))

	return nil // byeeee
}
