package others

import (
	"fmt"
	"sort"
)

func mostOccuringValue[T int | float32](slice []T) (T, T) {
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

type CountryInfo struct {
	Population  int
	Language    string
	CapitalCity string
}

func countries(countryName string, countryInfo CountryInfo) map[string]CountryInfo {
	countriesMap := make(map[string]CountryInfo)
	countriesMap[countryName] = CountryInfo{Population: countryInfo.Population, Language: countryInfo.Language, CapitalCity: countryInfo.CapitalCity}
	return countriesMap
}

func scores() {
	userScores := make(map[string]int)
	var name string
	var score int
	for {
		fmt.Print("Enter a name (or type 'done' to finish): ")
		fmt.Scanln(&name)
		if name == "" || name == "done" {
			break
		}
		fmt.Print("What is ", name, "'s score: ")
		fmt.Scanln(&score)
		userScores[name] = score
	}
	var sortedNames []string
	for name := range userScores {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)
	fmt.Println("\nNames in alphabetical order along with their scores:")
	for _, name := range sortedNames {
		fmt.Printf("%s: %d\n", name, userScores[name])
	}
	// return fmt.Sprintf("%s: %d\n", name, userScores[name])
}

func MainFuncForMaps() {
	arrForMap := []int{21, 13, 45, 21, 21, 56, 87, 66, 87, 53, 87, 99, 99, 99, 99, 99, 99, 99, 99}
	findMostOccuringValue, numOfTimesItOccured := mostOccuringValue(arrForMap)

	fmt.Println("most occuring value is:", findMostOccuringValue, "and it occured:", numOfTimesItOccured)

	nigeriaInfo := CountryInfo{
		Population:  300_000_000,
		Language:    "English",
		CapitalCity: "Abuja",
	}
	fmt.Println(countries("Nigeria", nigeriaInfo))
	scores()
}
