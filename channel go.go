package main

import (
	"fmt"
	"time"
)

func mainj() {
	channel := make(chan string, 1)
	go func(ch chan string) {
		ch <- "2"
		fmt.Println(1)
	}(channel)
	message := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
}
func main() {
	channel := make(chan string, 1)
	go func(ch chan string) {
		mess := <-ch
		ch <- "hey from anonymous"
		fmt.Println(mess)
		fmt.Println(1)

	}(channel)
	message := "Hello from MaAIN FUNCTION"
	channel <- message
	time.Sleep(time.Second * 5)
	message = <-channel
	fmt.Println("message")

}
