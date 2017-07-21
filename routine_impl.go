package main



import "fmt"
import "sync"
import "runtime"
//import "time"

 

func main(){
	fmt.Println("Go routine program")
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
	        defer wg.Done()

	         
		var i int
	       
        	for i=1;i<5;i++ {
			 //time.Sleep(2 * time.Second)
			fmt.Println(i)
		}
        }()

	go func() {
        	defer wg.Done()
		var i int
		 
		for i=1;i<5;i++ {
			// time.Sleep(1 * time.Second)
			fmt.Println(i+10)
		}
        
        }()

	fmt.Println("Waiting To Finish")
        wg.Wait()

        fmt.Println("\nTerminating Program")
}


