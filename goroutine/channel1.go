package main

import (
  "fmt"
  "time"
)

func pinger(c chan<- string) {
  for i := 0; ; i++ {
    c <- "ping"
  }
}

func ponger(c chan<- string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func printer(c <- chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
//  var c chan string = make(chan string)
  fmt.Println("This is befor go routine call")
  //go pinger(c)
 // go ponger(c)
 // go printer(c)
   go func() {
    fmt.Println("This is inside the go routine")
  }()

  fmt.Println("This is after the go routine call")
  var input string
  fmt.Scanln(&input)
}
