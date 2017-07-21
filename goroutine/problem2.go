package main


import (
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup
var m sync.Mutex
func extract(i int){
	m.Lock()			
	fmt.Println(i, "is extracting")
	time.Sleep(5* time.Second)
	fmt.Println(i, "is extracted")
	m.Unlock()
	wg.Done()
}


func main()   {

	var i int

	
	wg.Add(3)

        for i=0;i<3;i++ {
	fmt.Println(i, "is downloading")
	time.Sleep(2 * time.Second)
	fmt.Println(i, "is downloaded")	
	go extract(i)

	}
       
     wg.Wait()
     

 

}
