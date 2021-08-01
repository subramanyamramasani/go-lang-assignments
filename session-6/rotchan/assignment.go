package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func show(a chan int, value int) {
	defer wg.Done()
	a <- value
}
func main() {
	showvalue := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go show(showvalue, i)
	}

	wg.Wait()
	close(showvalue)

	for data := range showvalue {
		fmt.Println(data)
		time.Sleep(time.Millisecond * 10)
	}
}
func call(i int) {
	fmt.Println(i)
}
