package main
import (
    "fmt"
    "io"
)
func main() {
    r, w := io.Pipe()
    go func() {
        w.Write([]byte("hello"))
        w.Close()
        w.Write([]byte("world"))
    }()
    buf := make([]byte, 100)
    for i := 0; i < 3; i++ {
        n, err := r.Read(buf)
        fmt.Printf("%q %v\n", buf[:n], err)
    }
}
