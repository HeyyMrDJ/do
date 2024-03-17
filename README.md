# do
Simple task runner written in Go

## Getting Started
- Create a Dofile and create a function and commands
```c
test() {
    echo hello world
}
```

- Build and run
```bash
go build -o do main.go
./do test
Running function: test
Executing command: echo hello world
hello world
```
