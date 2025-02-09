```go
func main() {
    var m sync.Mutex
    var wg sync.WaitGroup
    ch := make(chan int)

    wg.Add(1)
go1:
    go func() {
        defer wg.Done()
        m.Lock()
        select {
        case <-ch:
            fmt.Println("Exiting goroutine")
            m.Unlock()
        default:
            fmt.Println("Waiting...")
            time.Sleep(1 * time.Second)
            goto go1
        }
    }()

    time.Sleep(2 * time.Second)
    close(ch)
    wg.Wait()
}
```