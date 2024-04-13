package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("the value is %v\n", i)
		}()
		fmt.Println("return from fucntion")
		return
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("Done....")

}
