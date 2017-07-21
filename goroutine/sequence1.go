package main

import( "fmt"
)

func main()  { 


        fmt.Println("This is inside the main")

        go func() {
                fmt.Println("This is inside the goroutine")
        }()

        fmt.Println("This is after the goroutine")
        
}


