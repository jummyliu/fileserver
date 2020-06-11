package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

var (
	host string
	port int
	dir string
	h bool
	addr string
)

func init() {
	flag.StringVar(&host, "addr", "0.0.0.0", "addr")
	flag.IntVar(&port, "port", 8000, "port")
	flag.StringVar(&dir, "path", ".", "basic path")
	flag.BoolVar(&h, "h", false, "help")

	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}
	addr = fmt.Sprintf("%s:%d", host, port)
}

// Response custom response type to save response code
type Response struct {
	http.ResponseWriter
	status int
}

// WriteHeader override WriteHeader to save response code
func (resp *Response) WriteHeader(code int) {
	resp.status = code
	resp.ResponseWriter.WriteHeader(code)
}

func fileServerHandler(resp http.ResponseWriter, req *http.Request) {
	filePath, _ := filepath.Abs(path.Join(dir, req.URL.Path))
	if file, err := os.Stat(filePath); err == nil && !file.IsDir() {
		// modify response header
		resp.Header().Add("Content-Disposition", "attachment;filename=" + file.Name())
	}
	w := &Response{
		ResponseWriter: resp,
	}
	http.ServeFile(w, req, filePath)
	// ip | code | method | url
	if w.status == 0 {
		w.status = http.StatusOK
	}
	// log
	log.Printf("%s\t| %d\t| %s\t| %s", req.RemoteAddr, w.status, req.Method, req.URL.Path)
}

func main() {
	log.Printf("listen on %v", addr)
	if err := http.ListenAndServe(addr, http.HandlerFunc(fileServerHandler)); err != nil {
		log.Fatalf("listen %v failure: %v", addr, err)
	}
}
