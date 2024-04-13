package main

import (
	"fmt"
	"sync"
)

func main() {
	var data int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()
	wg.Wait()
	fmt.Printf("the value is %v\n", data)
	fmt.Println("Done....")

}
