  package main

  import (
          "fmt"
          "time"
  )

  func delaySecond(n time.Duration) {
          time.Sleep(n * time.Second) // <------------ here
  }

  func delayMinute(n time.Duration) {
          time.Sleep(n * time.Minute)
  }

  func main() {

          fmt.Println("starts")

          delaySecond(3) // delay 3 seconds

          fmt.Println("after 3 second")

          delayMinute(1) // delay 1 minute

          fmt.Println("after 1 minute delay... now get back to work!")

  }
