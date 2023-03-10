package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go func() {
		channel <- 1
		time.Sleep(time.Second)
		channel <- 2
		close(channel)
	}()
	for Ch := range channel {
		fmt.Println(Ch)
	}
}
func main() {
	helloCh := make(chan string, 1)
	goodByeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go receiveMessage(helloCh, goodByeCh, quitCh)
	go sendMessage(helloCh, "Hello world")
	time.Sleep(time.Second)
	go sendMessage(goodByeCh, "Good Bye world")
	result := <-quitCh
	fmt.Println("message from quitCh=", result)
}
func sendMessage(Ch chan<- string, message string) {
	Ch <- message

}
func receiveMessage(helloCh, goodByeCh <-chan string, quitCh chan<- bool) {
	for {
		select {
		case message := <-helloCh:
			fmt.Println("Message from helloCh: ", message)
		case message := <-goodByeCh:
			fmt.Println("Message from goodByeCh: ", message)
		case <-time.After(time.Second * 2):
			fmt.Println("nohing receiving in 2 seconds: Exiting the receiveMessage goroutinge")
			quitCh <- true
			break
		}
	}
}
