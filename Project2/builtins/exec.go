package builtins

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	ErrCouldNotStartProgram = errors.New("exec could not start program")
	ErrTooFewArguments      = errors.New("too few arguments")
)

func Exec(exit chan<- struct{}, args ...string) error {

	if len(args) == 0 {
		return fmt.Errorf("Error: %w", ErrTooFewArguments)
	}

	realArgs := args[1:] // get the actual args without the command name

	// execute command
	cmd := exec.Command(args[0], realArgs...)
	err := cmd.Start()

	// if it fails, don't kill the shell
	if err != nil {
		return fmt.Errorf("Error: %w", ErrCouldNotStartProgram)
	}

	// shell is kill...
	exit <- struct{}{}
	return nil
}
