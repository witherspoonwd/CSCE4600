package builtins

import (
	"fmt"
	"os/exec"
)

func Exec(exit chan<- struct{}, args ...string) error {

	if len(args) == 0 {
		return nil // bye (DONT KILL! FOLLOWING UBUNTU EXEC PATTERN)
	}

	realArgs := args[1:] // get the actual args without the command name

	// execute command
	cmd := exec.Command(args[0], realArgs...)
	err := cmd.Start()

	// if it fails, don't kill the shell
	if err != nil {
		return fmt.Errorf("Error: ", err)
	}

	// shell is kill...
	exit <- struct{}{}
	return nil
}
