package main

import (
 // "bytes"
 // "compress/gzip"

 "github.com/klauspost/pgzip"
  "io/ioutil"
  "fmt"
  "os"
)

func main()  {

        var zipFile = "hello_world.txt.gz"
	
	handle, err := openFile(zipFile)
	if err != nil {
        	fmt.Println("[ERROR] Opening file:", err)
    	}

	zipReader, err := pgzip.NewReaderN(handle, 100000, 10)
    
	if err != nil {
	        fmt.Println("[ERROR] New gzip reader:", err)
	    }


	 defer zipReader.Close()

    	fileContents, err := ioutil.ReadAll(zipReader)
    
	if err != nil {
	        fmt.Println("[ERROR] ReadAll:", err)
	}

   	fmt.Printf("[INFO] Uncompressed contents: %s\n", fileContents)

	closeFile(handle)

	
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
