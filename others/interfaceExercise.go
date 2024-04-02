package others

import (
	"fmt"
	"math"
)

type Calculator interface {
	Add(a, b interface{}) interface{}
	Subtract(a, b interface{}) interface{}
	Multiply(a, b interface{}) interface{}
	Divide(a, b interface{}) interface{}
}

type IntegerCalculator struct{}

func (calculator *IntegerCalculator) Add(a, b interface{}) interface{} {
	aValue, ok := a.(int)
	if !ok {
		return nil
	}
	bValue, ok := b.(int)
	if !ok {
		return nil
	}
	return aValue + bValue
}

func (calculator *IntegerCalculator) Subtract(a, b interface{}) interface{} {
	return a.(int) - b.(int)
}

func (calculator *IntegerCalculator) Divide(a, b interface{}) interface{} {
	return a.(int) / b.(int)
}

func (calculator *IntegerCalculator) Multiply(a, b interface{}) interface{} {
	return a.(int) * b.(int)
}

type FloatCalculator struct {
	s string
	i int
}

func (calculator *FloatCalculator) Add(a, b interface{}) interface{} {
	fmt.Println(calculator.s)
	fmt.Println(calculator.i)
	return a.(float64) + b.(float64)
}

func (calculator *FloatCalculator) Subtract(a, b interface{}) interface{} {
	return a.(float64) - b.(float64)
}

func (calculator *FloatCalculator) Divide(a, b interface{}) interface{} {
	return math.Round(a.(float64) / b.(float64))
}

func (calculator *FloatCalculator) Multiply(a, b interface{}) interface{} {
	var aValue float64
	var bValue float64
	switch v := a.(type) {
	case float64:
		aValue = v
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
	switch v := b.(type) {
	case float64:
		bValue = v
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
	return aValue * bValue
}

func cal(c Calculator) *Calculator {
	return &c
}

func MainFuncForInterfaceExercise() {
	// Create instances of the calculators
	intCalc := IntegerCalculator{}
	floatCalc := FloatCalculator{
		s: "hello world",
	}

	// Test integer arithmetic
	fmt.Println("Integer Arithmetic:")
	fmt.Println("Addition:", intCalc.Add(10, 5))
	fmt.Println("Subtraction:", intCalc.Subtract(10, 5))
	fmt.Println("Multiplication:", intCalc.Multiply(10, 5))
	fmt.Println("Division:", intCalc.Divide(10, 5))

	// Test floating-point arithmetic
	fmt.Println("\nFloating-Point Arithmetic:")
	fmt.Println("Addition:", floatCalc.Add(10.5, 5.5))
	fmt.Println("Subtraction:", floatCalc.Subtract(10.5, 5.5))
	fmt.Println("Multiplication:", floatCalc.Multiply(10.5, 5.5))
	fmt.Println("Division:", floatCalc.Divide(10.5, 5.5))
	c := cal(&intCalc)
	fmt.Println(c)
}
