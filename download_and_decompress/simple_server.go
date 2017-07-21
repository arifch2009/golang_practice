package main

import (
	//"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func gzipHandler(rw http.ResponseWriter, r *http.Request) {

	//var zipFile = "hello_world.txt.gz"
	
	//handle, err := openFile(zipFile)
/*	
	if err != nil {
        	fmt.Println("[ERROR] Opening file:", err)
    	}
*/

	log.Printf("Recv'd request")
	defer r.Body.Close()
	
	gr, err := gzip.NewReader(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()
	plaintext, err := ioutil.ReadAll(gr)
	
	
	
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Sample: %s", string(plaintext[0:100]))
	}
	//closeFile(plaintext)

	rw.Write([]byte(fmt.Sprintf("Recv'd %d plaintext bytes", len(plaintext))))
}

func main() {
	http.HandleFunc("/home", gzipHandler)
	go http.ListenAndServe(":8080", nil)

	in, err := os.Open("/home/aahmed/go/2.ts")
	if err != nil {
		log.Fatal(err)
	}
	// gzip writes to pipe, http reads from pipe
	pr, pw := io.Pipe()

	// buffer readers from file, writes to pipe
	//bufin := bufio.NewReader(in)

	// gzip wraps buffer writer and wr
	gw := gzip.NewWriter(pw)

	// Actually start reading from the file and writing to gzip
	go func() {
		log.Printf("Start writing")
		//n, err := bufin.WriteTo(gw)
                n, err := io.Copy(gw, in)
		if err != nil {

			log.Fatal(err)

		}
		gw.Close()
		pw.Close()
                log.Printf("copied    %v %v", n, err)
	}()

	req, err := http.NewRequest("POST", "http://localhost:8080/home", pr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Making request")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Done! Received: \n%s", string(respBody))
}

func closeFile(handle *os.File) {
    if handle == nil {
        return
    }

    err := handle.Close()
    if err != nil {
        fmt.Println("[ERROR] Closing file:", err)
    }
}


func openFile(fileToOpen string) (*os.File, error) {
    return os.OpenFile(fileToOpen, openFileOptions, openFilePermissions)
}


const openFileOptions int = os.O_CREATE | os.O_RDWR
const openFilePermissions os.FileMode = 0660
