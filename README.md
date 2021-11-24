# Concurrent-API-Calls-with-Go
### What are GoRoutines?
A Goroutine is a lightweight thread managed by the Go runtime. It lives within the user thread space, scheduled by the language, and is managed entirely by the run-time system. The kernel has no knowledge of user-level threads and manages them as if a single-thread process. The overhead for assignment, suspension, resumption is much lower than OS threads. 
<br>
In JVMs, threads map directly 1:1 to OS threads which limits massice concurrency. OS threads have large fixed stack size and there is a cap on the number of them to run in a single VM due to the increasing memory overhead. 
<br>
In Go, stacks are dynamically sized
That is why you can run many millions more goroutines and only thousands of Java threads at a time.

<br>

### The program
 This is a simple program demonstrating how goroutines are used for making concurrent API requests. <br>
 New routines are spawn with the go keyword before the function call. Goroutines share data over a channel. This avoids the problems of deadlocks and race conditions that commonly occur in other programming languages when accessing shared memory.
 
Run the program by downloading the repository and run the following commands
```shell

go mod init async && go mod tidy && go run async.go
```
