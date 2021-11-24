# Concurrent-API-Calls-with-Go
A Goroutine is a lightweight thread managed by the Go runtime with its own routine scheduler managing its lifecycle. It exists in the virtual space of a go runtime rather than in the OS, living within the user thread space, so the overhead for assignment, suspension, resumption is much lower than OS threads.  Compared to threads in Java it's much more lightweight as Java threads map directly to OS threads. OS threads have large fixed stack size and there is a cap on the number of them to run in a single VM due to the increasing memory overhead.
<br> 
<br>
 This is a simple program demonstrating how goroutines are used for making concurrent API requests. <br>
 New routines are spawn with the go keyword before the function call. Goroutines share data over a channel. This avoids the problems of deadlocks and race conditions that commonly occur in other programming languages when accessing shared memory.
 
Run the program by downloading the repository and run the following commands
```shell

go mod init async && go mod tidy && go run async.go
```
