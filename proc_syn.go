
package main

import (
	"fmt"
	"time"
 
        "sync"
	"math/rand"

)


func download(layer string) {

	s1 := rand.NewSource(time.Now().UTC().UnixNano())
    	r1 := rand.New(s1)
	var t int
        t = r1.Intn(10)
	fmt.Println("I am inside", layer ,"and Let me sleep for ", t)
	time.Sleep(time.Duration(t)* time.Second)
}


func extract(layer string) {

        s1 := rand.NewSource(time.Now().UTC().UnixNano())
        r1 := rand.New(s1)

        var t int

	t = r1.Intn(10)
	fmt.Println("I am inside", layer ,"and Let me sleep for ", t)
	time.Sleep(time.Duration(t)* time.Second)


	
}

func main()  {
	
//	var i, j int
	var wg sync.WaitGroup
	fmt.Println("Starting Sync program")
	wg.Add(2)
       go func() {
        defer wg.Done()


	fmt.Println("Inside the 1st thread")
        time.Sleep(5 * time.Second)

	
	}()


	go func(){
		 defer wg.Done()

	        fmt.Println("Inside the 2nd thread")

	        time.Sleep(5 * time.Second)
		
		
	}()

	wg.Wait()
	fmt.Println("Downloading finished successfly!!")
}


