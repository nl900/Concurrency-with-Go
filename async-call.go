package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

const baseUrl = "https://github.com/"

func makeRequest(user string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(baseUrl + user)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- user + " Not found"
	} else {
		ch <- user + " Success"
	}

}

func AsyncRequest(users []string) {
	ch := make(chan string)
	var user string
	var wg sync.WaitGroup // track values expected from the channel
	for _, user = range users {
		wg.Add(1)                     // there is now 1 pending operation
		go makeRequest(user, ch, &wg) // launch a go routine per url
	}

	// close the channel in the background
	go func() {
		wg.Wait() // block the goroutine until WaitGroup counter is zero
		close(ch)
	}()
	// read from channel as they come in until it's closed
	for res := range ch {
		fmt.Println(res)
	}

}

func main() {
	users := []string{
		"torvalds",
		"gvanrossum",
		"poettering",
		"this should not exist",
	}
	AsyncRequest(users)
}
