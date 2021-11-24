# Concurrent-API-Calls-with-Go
A Goroutine is a function which executes independently and simultaneously in connection with other Goroutines present in your program. It is a lightweight thread managed by the Go runtime.
<br>
<br>
Goroutines exist only in the virtual space of a go runtime instead of in the OS and uses its own routine scheduler to manage the goroutine lifecycle. Goroutines live within the user thread space and so the overhead for assignment, suspension, resumption is much lower than OS threads. That is why it is much more lightweight than threads in Java which map directly to OS threads. OS threads have large fixed stack size and there is a cap on the number of them to run in a single VM due to the increasing memory overhead.
<br> 
<br>
 This is a simple program demonstrating how goroutines are used for making concurrent API requests. <br>
 New routines are spawn with the go keyword before the function call. Goroutines share data over a channel. This avoids the problems of deadlocks and race conditions that commonly occur in other programming languages when accessing shared memory.
 
Run the program by downloading the repository and run the following commands
```shell

go mod tidy && go run async.go
```
