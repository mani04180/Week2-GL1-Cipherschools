package main

import (
	"fmt"
	"time"
)

func main3() {

	current_time := time.Now()
	fmt.Println("ANSIC: ", current_time.Format(time.ANSIC))
	fmt.Println("UnixDate: ", current_time.Format(time.UnixDate))
	fmt.println("RFC1123: ", current_time.Format(time.RFC1123))
	fmt.Println("RFC3339Nano: ", current_time.Format(time.RFC3339Nano))
	fmt.Println("RubyDate: ", current_time.Format(time.RubyDate))
}
