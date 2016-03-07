package main

import (
	"fmt"
	"time"
)

func main() {

	var toto int = 42

	c := make(chan int)

	for i := 0; i < 4; i++ {
		go func(delay int) {
			time.Sleep(time.Duration(delay) * time.Second)
			fmt.Printf("Hello after %d sec\n", delay)

			c <- 0
		}(i)
	}
	for i := 0; i < 4; i++ {
		<-c
	}
	fmt.Println("Finished")
}
