package others

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure_area(g geometry) float64 {
	return g.area()
}

func extract_value(v interface{}) string {
	if number, ok := v.(int); ok {
		return fmt.Sprintf("Extracted value %v", number)
	} else {
		return fmt.Sprintf("Value not an integer %v", v)
	}
}

func MainFuncForInterfaces() {
	circle1 := Circle{
		radius: 12.3,
	}
	areaOfCircle := measure_area(&circle1)
	fmt.Println(math.Round(areaOfCircle))

	fmt.Println(extract_value(21))
}
