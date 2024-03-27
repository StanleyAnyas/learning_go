package others

import (
	"fmt"
	// "math/rand"
	"sync"
)

func putInChannel(nums []int) <-chan int {
	out := make(chan int, len(nums))

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(nums <-chan int) <-chan int {
	out := make(chan int, len(nums))

	go func() {
		for n := range nums {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func MainForOthers() {
	var wg sync.WaitGroup
	// wgs := sync.WaitGroup{} // you can either use this way or the one with va to initialize the sync.WaitGroup
	// mutex := sync.Mutex{}
	// mutex.Lock()
	// mutex.Unlock()

	nums := []int{2, 6, 9, 4, 3}

	// first stage out the slice into a go routine
	firstStage := putInChannel(nums)

	// second stage: square the numbers
	finalStage := sq(firstStage)

	// final stage: print the numbers
	for n := range finalStage {
		fmt.Println(n)
	}

	wg.Add(1) // add this before every goroutine so to increment the number of goroutines the main function has to wait before exiting
	go func() {
		defer wg.Done() // and always add this either at the end inside the goroutine function: wg.Done() or at the top of with but with defer keyword: defer wg.Done()
		for i := 0; i < 12; i++ {
			fmt.Println(i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 13; i < 26; i++ {
			fmt.Println(i)
		}

	}()

	wg.Wait() // this will make the main function to wait for the goroutines to finish

	fromExercise := mainForExercise()
	fmt.Println(fromExercise)

	floatSlice := []float32{7.2, 87.5, 23.7}
	fmt.Println(sumSlice(floatSlice))

}
