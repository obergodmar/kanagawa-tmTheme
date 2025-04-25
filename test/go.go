// Sample Go file
package main

import "fmt"

func greet(name string) {
    if name != "" {
        fmt.Println("Hello,", name)
    } else {
        fmt.Println("Hello, world!")
    }
}

func main() {
    greet("Developer")
}
