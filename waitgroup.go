package main

import "fmt"
import "time"
import "sync"


func main() {
     var wg sync.WaitGroup
     var i int
     wg.Add(3)
     fmt.Println("Before the main")
     for i=0;i<3;i++{ 
     go func() {
         defer wg.Done()
         time.Sleep(2* time.Second)
         fmt.Println("Inside the go routuine")
      }()
     }
     fmt.Println("After the main")

    wg.Wait()

}
