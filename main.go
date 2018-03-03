package main

import (
    "fmt"
    "os"
    "os/exec"
)

//go run main.go <cmd> <args>
func main() {
}

func run() {
    fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

    cmd := exec.Command()
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    must(cmd.Run())
}

func must(err error) {
    if err != nil {
        panic(err)
    }
}
