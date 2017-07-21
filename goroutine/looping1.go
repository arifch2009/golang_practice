package main

import (
	"fmt"
	"time"
)


func main() {
	
	
	for {
	fmt.Println("Let me sleep")
	time.Sleep(3 * time.Second)
	fmt.Println("Just woke up")
	break
	}

}
