package others

import (
	"fmt"
	"math/rand"
	// "sync"
)

// func isPrime(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}
// 	if num <= 3 {
// 		return true
// 	}
// 	if num%2 == 0 || num%3 == 0 {
// 		return false
// 	}
// 	i := 5
// 	for i*i <= num {
// 		if num%i == 0 || num%(i+2) == 0 {
// 			return false
// 		}
// 		i += 6
// 	}
// 	return true
// }

func generateRandomNum[K any, T any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func performOperation[K any](done <-chan K, num <-chan int, n int) <-chan int {
	operatedNums := make(chan int)

	go func() {
		defer close(operatedNums)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case operatedNums <- <-num:
			}
		}
	}()

	return operatedNums
}

func returnTwoValues(num int) (int, error) {
	return num * 2, nil
}

func mainForExercise() []int {
	arr := []int{}
	done := make(chan struct{}, 1)
	defer close(done)
	randNumFetcher := func() int { return rand.Intn(100000) }
	randomNum := generateRandomNum(done, randNumFetcher)

	for finalized := range performOperation(done, randomNum, 10) {
		arr = append(arr, finalized)
	}

	if num, err := returnTwoValues(2); err == nil {
		fmt.Println("returnTwoValues returned:", num)
	}

	return arr
}
