package builtins

import (
	"bufio"
	"fmt"
	"os"
)

func WriteHistory(input string) error {
	// open file
	file, err := os.OpenFile("history.dat", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Error: Unable to open history.dat. Shell history will not be saved this session. Message:", err)
	}
	defer file.Close()

	// write the command typed to the history
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(input) // Append content with a newline
	if err != nil {
		return fmt.Errorf("Error: Unable to write to history.dat. This command will NOT be saved in history! Message:", err)
	}

	// close write
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("Error: ", err)
	}

	return nil
}

func ClearHistory() error {
	file, err := os.OpenFile("history.dat", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Error: Unable to clear history.dat! Message:", err)
	}
	defer file.Close()

	return nil
}

func PrintHistory() error {
	// open file
	file, err := os.Open("history.dat")
	if err != nil {
		return fmt.Errorf("Error: Unable to clear history.dat! Message:", err)
	}
	defer file.Close()

	// make scanner
	scanner := bufio.NewScanner(file)

	// print each line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// if theres an error (RIP)
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error: Unable to clear history.dat! Message:", err)
	}

	return nil
}

func HistoryCommand(args ...string) error {
	// supports -c
	switch len(args) {
	case 1:
		if args[0] == "-c" {
			return ClearHistory()
		}
	default:
		return PrintHistory()
	}

	return nil
}
