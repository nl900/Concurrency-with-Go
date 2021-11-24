# Concurrent-API-Calls-with-Go
### Concurrency
Doing multiple tasks simultaneously. Go was designed to run on multiple cores and it's built to support concurrency and scale. It achieves this through GoRoutines.

### What are GoRoutines?
A Goroutine is a lightweight thread managed by the Go runtime. It lives within the user thread space, scheduled by the language, and is managed entirely by the run-time system. The kernel has no knowledge of user-level threads and manages them as if a single-thread process. So the overhead for assignment, suspension, resumption is much lower than OS threads. 
<br>
In JVMs, threads map directly 1:1 to OS threads which limits massive concurrency. OS threads have large fixed stack size, typically defaulting to 1MB per thread.  there is a cap on the number of them to run in a single VM due to the increasing memory overhead. 
<br>
In Go, stacks are dynamically sized. A new goroutine typically have an initial stack size of 4KB.
That is why you can run many millions more goroutines and only thousands of Java threads at a time.
<br>
A second issue is JVM relies on the OS kernel to schedule OS threads. The OS keeps a list of all running tasks, as  Linux doesn't distinguish between threads and processes, and attempts to give each a fair share of CPU time. When it switches from one thread to another, the new thread running must be started with a view of the world that abstract away other threads running on the same CPU. This is expensive.
<br>
Go has its own scheduler that allows many Goroutines to run on the same OS thread. It saves a significant amount of time on context switching. The scheduler also optimises by only running a Goroutine that has a non-empty channel with work to do.


### The program
 This is a simple program demonstrating how goroutines are used for making concurrent API requests. <br>
 New routines are spawn with the go keyword before the function call. Goroutines share data over a channel. This avoids the problems of deadlocks and race conditions that commonly occur in other programming languages when accessing shared memory.
 
Run the program by downloading the repository and run the following commands
```shell
$ go mod init async && go mod tidy && go run async.go
```
