package main

import (
	"fmt"
	"sync"
)

func main() {

	var wait sync.WaitGroup
	counter := 20000
	wait.Add(counter)
	for i := 0; i < counter; i++ {
		go hello(&wait, i)
	}
	defer wait.Wait()
}
func hello(wait *sync.WaitGroup, counter int) {
	fmt.Println(counter)
	wait.Done()
}
