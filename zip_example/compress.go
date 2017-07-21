package main

import (
  "bytes"
  //"compress/gzip"
  "github.com/klauspost/pgzip"
  "io/ioutil"
  "fmt"
)

func main()  {
	var b bytes.Buffer
	w := pgzip.NewWriter(&b)
	w.SetConcurrency(100000, 10)
	w.Write([]byte("hello, world\n"))
	w.Close() // You must close this first to flush the bytes to the buffer.
	err := ioutil.WriteFile("hello_world.txt.gz", b.Bytes(), 0666)
	fmt.Println(err)
}
