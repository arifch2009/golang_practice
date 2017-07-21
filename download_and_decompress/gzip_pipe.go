
package main


import (
	"fmt"
	"io"
	"os"
//	"gzip"

	"github.com/klauspost/pgzip"
)


func main() {
	
	var gzFilePath string ="/home/aahmed/go/download_and_decompress/hello_world.txt.gz"
	var dstFilePath string ="/home/aahmed/go/download_and_decompress/"
	UnpackGzipFile(gzFilePath, dstFilePath)

}


func UnpackGzipFile(gzFilePath, dstFilePath string) (int64, error) {
    gzFile, err := os.Open(gzFilePath)
    if err != nil {
        return 0, fmt.Errorf("Failed to open file %s for unpack: %s", gzFilePath, err)
    }
    dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
    if err != nil {
        return 0, fmt.Errorf("Failed to create destination file %s for unpack: %s", dstFilePath, err)
    }

    ioReader, ioWriter := io.Pipe()

    go func() { // goroutine leak is possible here
        //gzReader, _ := gzip.NewReader(gzFile)
        gzReader, _ := pgzip.NewReaderN(gzFile, 250000, 1)
	// it is important to close the writer or reading from the other end of the
        // pipe or io.copy() will never finish
        defer func(){
            gzFile.Close()
            gzReader.Close()
            ioWriter.Close()
        }()

        io.Copy(ioWriter, gzReader)
    }()

    written, err := io.Copy(dstFile, ioReader)
    if err != nil {
        return 0, err // goroutine leak is possible here
    }
    ioReader.Close()
    dstFile.Close()

    return written, nil
}
