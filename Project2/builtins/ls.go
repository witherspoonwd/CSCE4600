package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrCouldNotOpenDirectory = errors.New("couldnt open directory")
	OsDirectory, _           = os.Getwd()
)

func ListDirectory(args ...string) error {

	var dir string
	var err error

	switch len(args) {
	case 1:
		dir = args[0] // use what the user said
	default:
		dir = OsDirectory // get the os Directory
	}

	files, err := os.ReadDir(dir) //get files array and update err

	if err != nil {
		return fmt.Errorf("Error: %w", ErrCouldNotOpenDirectory)
	}

	// get max length for formatting
	maxLen := 0
	for _, file := range files {
		length := len(file.Name())
		if length > maxLen && maxLen < 25 {
			maxLen = length
		}
	}

	// print the files
	for i, file := range files {
		name := file.Name()
		fmt.Printf("%-*s", maxLen, name)

		if (i+1)%6 == 0 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")

	return nil // byeeee
}
