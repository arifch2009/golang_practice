package main

import (
  "fmt"
  "time"
//   "strconv"
   "math/rand"
)

var j int = 0

func radom_gen() int {
	s1 := rand.NewSource(time.Now().UTC().UnixNano())
    	r1 := rand.New(s1)
	return r1.Intn(10)

}
func downloader(layer int,  c chan int) {
	
	fmt.Println(layer, "is downloading")
	
	time.Sleep(time.Second * time.Duration(radom_gen()) )
  	fmt.Println(layer, "is downloaded")
	c <- layer
	
}

func extractor( c chan int) {
    
  msg := <- c
   if msg==0 && j==0  {  
     
     
    fmt.Println(msg, "is extracting")
    time.Sleep(time.Second * time.Duration(radom_gen()))
    fmt.Println(msg,  "is extracted")
     j=1
   } else if msg==1  && j==1  {  
      
     
    fmt.Println(msg, "is extracting")
    time.Sleep(time.Second * time.Duration(radom_gen()))
    fmt.Println(msg,  "is extracted")
     j=2
   }   else if msg==2  && j==2  {  

    
    fmt.Println(msg, "is extracting")
    time.Sleep(time.Second * time.Duration(radom_gen()))
    fmt.Println(msg,  "is extracted")
     j=2
   }   else {
  	fmt.Println("Esle statement")
    }

}

func main() {
  var c chan int = make(chan int)
  var i int
  for i=0;i<3;i++ {
  go downloader(i, c)
  }
  for i=0;i<3;i++ {  
  go extractor( c)
  }
  var input string
 fmt.Scanln(&input)
}
