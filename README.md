# Shell Project

## implemented builtins

### history

- automatically stores history of all command attempts.
- supports clearing with "history -c"
- history prints history

usage:

```
  history [arguments]
```

### ls

- prints the current working directory
- can take a filepath as a directory to print it

```
  ls
```

### exec

- executes a program while killing the shell

```
  exec [command]
```

### time

- executes a program while displaying the amount of time it took for the program to run

```
  time [command]
```

### echo

- echos arguments passed to it

```
  echo [thing to echo]
```

## compilation instructions

1. clone this repository
2. run ```go build main.go``` in "CSCE4600/Project2/"
3. run the program

## testing instructions
1. navigate to ```./Project2```
2. run ```golangci-lint run```
3. test