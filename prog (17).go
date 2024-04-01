package main

import (
	"fmt"
    "os"
    "io"
)
// to create any text file in folder access the file run and show in terminal
func main() {
	f,err:=os.Open(os.Args[1])
    if err!=nil{
        fmt.Println("Error",err)
        os.Exit(1)
    }
    io.Copy(os.Stdout,f)
}
