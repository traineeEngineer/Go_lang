package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}
func main() {
	fun("direct call")

	go fun("routine-2")

	go func() {
		fun("routine 3")
	}()

	fv := fun
	go fv("routine4")

	fmt.Println("wait for Go routine")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done.....")
}
