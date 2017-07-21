	

package main


import (
        "fmt"
      "sync"
        "time"
)


var k int

var m sync.Mutex

func download(i int){

       //print_something_here(i)
		
        go func() {
       
	 
	for  k != i{
		  
	}

	 
	print_something_here(i)
	k=k+1		 

	 
        }()

}

func print_something_here(i int){
	
        fmt.Println("Printing Starts here", i)

        time.Sleep(time.Second * 2)
        fmt.Println("Printing Ends here", i)
	



}



func main ()  {
	var i int
	//a[0]=0
	//fmt.Println("Length of a :", len(a))
        for i=0;i<4;i++ {
        	
		download(i)
        }
	var input string
  	fmt.Scanln(&input)

 }



