package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	a := 10
	b := 30

	defer fmt.Println("value of a is:", a)
	defer fmt.Println("value of b is:", b)

	fmt.Println("processor :", runtime.NumCPU())
	fmt.Println("Os :", runtime.GOOS)
	fmt.Println("Architecture :", runtime.GOARCH)
	fmt.Println("Max processes :", runtime.GOMAXPROCS(8))

	wg.Add(2)
	go Positive(&wg)
	go Negative(&wg)

	fmt.Println("invoked Routine")

	wg.Wait()
}

func Positive(wg *sync.WaitGroup) error {
	fmt.Println("positive nums")
	numrange := 20
	for i := 0; i <= numrange; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	wg.Done()
	return nil
}

func Negative(wg *sync.WaitGroup) error {
	fmt.Println("negative nums")
	numrange := 20
	for i := 0; i <= numrange; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
	wg.Done()
	return nil
}
