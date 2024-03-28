package others

// import (
// 	"maps"
// )

func MostOccuringValue[T int | float32](slice []T) (T, T) {
	myMap := make(map[T]T)
	for _, n := range slice {
		myMap[n]++
	}

	var maxCount T
	var mostOccuringValue T

	for count, value := range myMap {
		if count > maxCount {
			maxCount = count
			mostOccuringValue = value
		}
	}

	// fmt.Println("mostOccuringValue:", maxCount, "and it occured:", mostOccuringValue)
	return maxCount, mostOccuringValue
}
