# Concurrent-API-Calls-with-Go
A Goroutine is a function which executes independently and simultaneously in connection with other Goroutines present in your program. Compared to threads, 
Goroutines exist only in the virtual space of a go runtime instead of in the OS. Go uses its own routines scheduler to manage the goroutine lifecycle. 
<br>

 <br>
 
 This is a simple program demonstrating how goroutines are used for making concurrent API requests. <br>
 New routines are spawn with the go keyword before the function call. Goroutines share data over a channel. This avoids the problems of deadlocks and race conditions that commonly occur in other programming languages.
