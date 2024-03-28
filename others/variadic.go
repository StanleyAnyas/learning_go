package others

import "fmt"

func Variadic_func(nums ...int) string {
	total := 0

	for _, num := range nums {
		total += num
	}
	result := fmt.Sprintf("The total sum of %v is %v", nums, total)
	return result
}

func Closure() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func factorial(nums int) int {
	if nums == 0 {
		return 1
	}
	return nums * factorial(nums-1)
}

func MainFuncForVariadic() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(Variadic_func(nums...))

	counter := Closure()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println(factorial(12))
}
