package main

import (
	"fmt"
	"time"
)

func doSomething(s string) {
  fmt.Println("doing something", s)
}

func main() {
	
  for {
    <-time.After(2 * time.Second)
    go doSomething("from polling 2")
     }

}

