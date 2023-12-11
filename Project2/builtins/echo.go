package builtins

import (
	"errors"
	"fmt"
)

var (
	ErrNoWriteToShell = errors.New("could not write to shell")
)

func Echo(args ...string) error {

	// print the stuff
	for _, arg := range args {
		_, err := fmt.Printf("%s ", arg)

		if err != nil {
			return fmt.Errorf("%w Unable to write to shell. Echo failed", ErrNoWriteToShell)
		}
	}

	fmt.Printf("\n")

	return nil // byeeee
}
