package main

import (
  "fmt"
  "time"
   "strconv"

)

func downloader(layer int,  c chan string) {
	
	fmt.Println(layer, "is downloading")
	time.Sleep(time.Second * 2)
  	fmt.Println(layer, "is downloaded")
	c <- "layer"+ strconv.Itoa(layer)
	
}

func extractor( c chan string) {

    msg := <- c
    fmt.Println(msg, "is extracting")
    time.Sleep(time.Second * 2)
    fmt.Println(msg,  "is extracted")
  
}

func main() {
  var c chan string = make(chan string)
  var i int
  for i=0;i<3;i++ {
  go downloader(i, c)
  }
  for i=0;i<3;i++ { 

  go extractor(c)
  }
   var input string
  fmt.Scanln(&input)
 
}
