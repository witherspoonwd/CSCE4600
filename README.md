# Shell Project

## implemented builtins

### history

- automatically stores history of all command attempts.
- supports clearing with "history -c"
- history prints history

### ls

- prints the current working directory
- can take a filepath as a directory to print it

### exec

- executes a program while killing the shell

### time

- executes a program while displaying the amount of time it took for the program to run

### echo

- echos arguments passed to it

## compilation instructions

1. clone this repository
2. run "go build main.go" in "$DIRECTORY$/Project2/"
3. run the program