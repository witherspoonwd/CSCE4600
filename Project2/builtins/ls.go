package builtins

import (
	"fmt"
	"os"
)

func ListDirectory(args ...string) error {

	var dir string
	var err error

	switch len(args) {
	case 1:
		dir = args[0] // use what the user said
		break
	default:
		dir, err = os.Getwd() // get the directory were in
		if err != nil {
			return fmt.Errorf("Error: ", err)
		}
	}

	files, err := os.ReadDir(dir) //get files array and update err

	if err != nil {
		return fmt.Errorf("Error:", err)
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
