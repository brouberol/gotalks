package main

import (
	"fmt"
	"math/rand"
	"time"
)

func HelloWorld(i int, c chan string) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	c <- fmt.Sprintf("%d - Hello World\n", i)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // initiate the seed
	c := make(chan string)
	for i := 1; i <= 10; i++ {
		go HelloWorld(i, c)
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf(<-c) // block until the channel is ready to send a value, and print it
	}
	close(c)
}
