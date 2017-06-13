package main

import (
    "os"
    "fmt"
)
func main() {
    buf := []byte{}
    n, _ := os.Stdin.Read(buf)
    fmt.Println("Read", n, "lines")
    fmt.Println("Buf", string(buf))
}
