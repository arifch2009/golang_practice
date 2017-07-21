package main

import (
  "io/ioutil"; 
  )


func main() {
  contents,_ := ioutil.ReadFile("origin.txt")
  println(string(contents))
  ioutil.WriteFile("destination", contents, 0644)
}
