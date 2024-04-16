package others

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
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

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done working...")
	done <- true
}

func writeToBufferChan(toBeReturned chan string, word string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, char := range word {
		toBeReturned <- string(char)
	}
	close(toBeReturned)
	// return toBeReturned
}

func f(from string, channel chan<- int, wg ...*sync.WaitGroup) {
	if wg != nil {
		defer wg[0].Done()
	}

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
	channel <- 2
	// return channel
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(ping <-chan string, pongs chan<- string) {
	msg := <-ping
	pongs <- msg
}

func MainForOthers() {
	var wg sync.WaitGroup
	// wgs := sync.WaitGroup{} // you can either use this way or the one with va to initialize the sync.WaitGroup
	// mutex := sync.Mutex{}
	// mutex.Lock()
	// mutex.Unlock()

	message := make(chan int)

	// f("workings")
	// var wg sync.WaitGroup
	wg.Add(1)
	go f("goroutine", message, &wg)

	wg.Add(1)
	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("going")

	fmt.Println(message)
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

	wg.Add(1)
	word := "Hello, world!"
	chanToWriteBuffer := make(chan string, len(word))

	go writeToBufferChan(chanToWriteBuffer, word, &wg)
	for data := range chanToWriteBuffer {
		fmt.Println("data:", data)
	}
	// fmt.Println("chanToWriteBuffer ", chanToWriteBuffer)

	fromExercise := mainForExercise()
	fmt.Println(fromExercise)

	floatSlice := []float32{7.2, 87.5, 23.7}
	fmt.Println(sumSlice(floatSlice))

	done := make(chan bool, 1)
	go worker(done)
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message passed")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	// wg.Wait() // this will make the main function to wait for the goroutines to finish

}
