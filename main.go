package main

import (
    "fmt"
    "sync"

)

func main() {
    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func(n int) {
            defer wg.Done()
            fmt.Println("Goroutine", n)
        }(i)
    }
    wg.Wait()
    fmt.Println("All goroutines finished executing")
}
