package builtins

import (
	"fmt"
)

func Echo(args ...string) error {

	// print the stuff
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}

	fmt.Printf("\n")

	return nil // byeeee
}
