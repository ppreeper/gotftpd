package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/pin/tftp"
)

var (
	addr string
	path string
)

func init() {
	flag.StringVar(&addr, "addr", "0.0.0.0:69", "Address to listen")
	flag.StringVar(&path, "path", ".", "Local file path")
	flag.Parse()
}

func main() {
	//use nil in place of handler to disable read and write operations
	s := tftp.NewServer(readHandler, writeHandler)
	s.SetTimeout(5 * time.Second) // optional
	err := s.ListenAndServe(addr) // blocks until s.Shutdown() is called
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {

	file, err := os.Open(filepath.Join(path, filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	//Set transfer size before calling ReadFrom
	rf.(tftp.OutgoingTransfer).SetSize(fi.Size())

	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)
	return nil
}

// writeHandler is called when client starts file upload to server
func writeHandler(filename string, wt io.WriterTo) error {
	file, err := os.OpenFile(filepath.Join(path, filename), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	defer file.Close()

	n, err := wt.WriteTo(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes received\n", n)
	return nil
}
