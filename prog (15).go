package main

import (
	"fmt"
	"net/http"
	"os"
)
type logWriter struct{}


func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
    lw:=logWriter{}

	io.Copy(lw,resp.body)
}

func (logWriter) Write(bs []byte)(int,error){
    fmt.Println((string(bs)))
    fmt.Println("just wrote  this many byte",len(bs))
    return len(bs), nil
}
