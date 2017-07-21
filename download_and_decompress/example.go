// Pipe creates a synchronous in-memory pipe.
// It can be used to connect code expecting an io.Reader
// with code expecting an io.Writer.
// Reads on one end are matched with writes on the other,
// copying data directly between the two; there is no internal buffering.
// It is safe to call Read and Write in parallel with each other or with
// Close. Close will complete once pending I/O is done. Parallel calls to
// Read, and parallel calls to Write, are also safe:
// the individual calls will be gated sequentially.

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, w := io.Pipe()

	go func() {
		defer w.Close()
		json.NewEncoder(w).Encode(map[string]string{
			"Name": "Mohanson",
		})
	}()

	resp, err := http.Post("https://httpbin.org/post", "application/json", r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
}
