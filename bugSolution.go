```go
package main

import (

    "fmt"
    "sync"
    "time"
)

func main() {
    var m sync.Mutex
    var wg sync.WaitGroup
    ch := make(chan int)

    wg.Add(1)
    go func() {
        defer wg.Done()
        m.Lock()
        for {
            select {
            case <-ch:
                fmt.Println("Exiting goroutine")
                m.Unlock()
                return
            default:
                fmt.Println("Waiting...")
                time.Sleep(1 * time.Second)
            }
        }
    }()

    time.Sleep(2 * time.Second)
    close(ch)
    wg.Wait()
}
```