package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

// FunctionBlock represents a block of commands under a function name
type FunctionBlock struct {
    Name     string
    Commands []string
}

func openDofile() ([]FunctionBlock, error) {
    file, err := os.Open("./Dofile")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var blocks []FunctionBlock
    var currentBlock *FunctionBlock
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        line = strings.TrimSpace(line)

        // Detect the start of a function block and strip off the parentheses
        if strings.HasSuffix(line, "{") {
            blockNameWithParentheses := strings.TrimSpace(strings.TrimRight(line, "{"))
            blockName := strings.TrimSpace(strings.TrimSuffix(blockNameWithParentheses, "()"))
            currentBlock = &FunctionBlock{Name: blockName}
            continue
        }

        // Detect the end of a function block
        if line == "}" && currentBlock != nil {
            blocks = append(blocks, *currentBlock)
            currentBlock = nil
            continue
        }

        // If inside a function block, add commands to the current block
        if currentBlock != nil {
            currentBlock.Commands = append(currentBlock.Commands, line)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return blocks, nil
}

func executeCommands(commands []string) {
    for _, cmd := range commands {
        fmt.Println("Executing command:", cmd)
        // Execute the command using bash -c
        command := exec.Command("bash", "-c", cmd)
        output, err := command.CombinedOutput()
        if err != nil {
            fmt.Println("Error executing command:", err)
            continue
        }
        fmt.Println(string(output))
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <function_name>")
        return
    }
    functionName := os.Args[1]

    blocks, err := openDofile()
    if err != nil {
        fmt.Println("Error parsing Dofile:", err)
        return
    }

    for _, block := range blocks {
        if block.Name == functionName {
            fmt.Printf("Running function: %s\n", functionName)
            executeCommands(block.Commands)
            return
        }
    }

    fmt.Printf("Function '%s' not found\n", functionName)
}
