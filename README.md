# Concurrent-API-Calls-with-Go

Go uses its own routines scheduler to abstract its lightweight goroutines from OS threads. 
<br>
Go always has at least one goroutine running, which takes care of running main(). New routines are spawn with the go keyword before the function call. 
