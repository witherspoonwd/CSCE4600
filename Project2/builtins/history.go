package builtins

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var (
	HistoryPath      string = "history.dat"
	ErrNoOpenHistory        = errors.New("could not open history.dat")
)

func WriteHistory(input string) error {
	// open file
	file, err := os.OpenFile(HistoryPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Error: Unable to open history.dat: %w", ErrNoOpenHistory)
	}
	defer file.Close()

	// write the command typed to the history
	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(input) // Append content with a newline
	if err != nil {
		return fmt.Errorf("Error: Unable to open history.dat: %w", ErrNoOpenHistory)
	}

	// close write
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("Error: Unable to open history.dat: %w", ErrNoOpenHistory)
	}

	return nil
}

func ClearHistory() error {
	file, err := os.OpenFile(HistoryPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Error: Unable to open history.dat: %w", ErrNoOpenHistory)
	}
	defer file.Close()

	return nil
}

func PrintHistory() error {
	// open file
	file, err := os.Open(HistoryPath)
	if err != nil {
		return fmt.Errorf("Error: Unable to open history.dat: %w", ErrNoOpenHistory)
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
		return fmt.Errorf("Error: Unable to clear history.dat! Message: %w", err)
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
