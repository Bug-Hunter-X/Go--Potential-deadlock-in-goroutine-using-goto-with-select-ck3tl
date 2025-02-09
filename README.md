# Go: Potential deadlock in goroutine using goto with select

This repository demonstrates a potential deadlock scenario in Go. A goroutine uses `goto` to repeatedly check a channel. If the channel is closed before the goroutine receives a value, it can deadlock.

## Bug Description

A goroutine continuously checks a channel using `goto` and a `select` statement with a `default` case. If the channel is closed before the goroutine receives a value, the `default` case is always executed, leading to an infinite loop and a deadlock.

## Bug Reproduction

1. Clone the repository.
2. Run `go run bug.go`.
3. Observe the potential deadlock (program might hang).

## Solution

The solution involves removing the `goto` and using a `for` loop with a condition to check the channel's state. This approach allows for graceful termination without causing a deadlock.

## Additional Notes

- This example highlights the potential dangers of using `goto` in concurrency scenarios in Go.
- It's crucial to carefully manage channels and goroutines to avoid deadlocks and other concurrency issues.