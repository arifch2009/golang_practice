package main 


import (  
	"fmt"
	"time"
)

func main()  {
	var i int
	var k *int
	var l int
	for l=0;l<5;l++ {
	go func() {
	k=&i
	fmt.Println("Intitial value of i is :", i, *k)
	*k=*k+1
	fmt.Println("After the increament i is: ", i , *k)
	}()
	}

	time.Sleep( 2 * time.Second)	
}

