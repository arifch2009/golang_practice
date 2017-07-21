package main

import (
  "fmt"
  "time"
//   "strconv"
//   "sync"
   "math/rand"
   
)


func radom_gen() int {
	s1 := rand.NewSource(time.Now().UTC().UnixNano())
    	r1 := rand.New(s1)
	return r1.Intn(10)

}


func func1( c <- chan int ) {
    
     // var c1 chan int = make(chan int)
    msg :=  <-c    
    if msg==0 {
//    m.Lock()
    fmt.Println(msg, "is extracting")
    time.Sleep(time.Second * 5)
    fmt.Println(msg,  "is extracted")
  //  m.Unlock()
    go func2(c)
    
    } 
  

}


func func2 ( c <- chan int) {
    
     // var c2 chan int = make(chan int)
    msg :=  <-c    
    if msg==1 {
   // m.Lock()
   fmt.Println(msg, "is extracting")
    time.Sleep(time.Second * 5)
    fmt.Println(msg,  "is extracted")
  //  m.Unlock()
    go func3(c)      
    }


}

  
func func3 ( c <- chan int) {
      
    
    msg :=  <-c    
    if msg==2 {
   // m.Lock()
    fmt.Println(msg, "is extracting")
    time.Sleep(time.Second * 5)
    fmt.Println(msg,  "is extracted")
  //  m.Unlock()
      
    }

}


func download (i int, c chan int) {

        fmt.Println(i, "is downloading")
       //var t int
       //t= radom_gen()
        time.Sleep(time.Second * 2)


        fmt.Println(i, "is downloaded and took", 2)
 
        c <- i

     }


  var c chan int = make(chan int)

func main() {

    
//   m := new(sync.Mutex)
  var i int

  //var c chan int = make(chan int)
  
  for i=0;i<3;i++ {
     go download(i,c)
  }

  go func1(c)


   
    

   var input string
  fmt.Scanln(&input)
 
}
