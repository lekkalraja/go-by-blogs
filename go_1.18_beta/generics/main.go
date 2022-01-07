package main

import "fmt"

// SumInts sums all the int64 values of map and returns the total sum : NON-GENERIC FUNC
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumInts sums all the float64 values of map and returns the total sum : NON-GENERIC FUNC
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumOfNumbers is a GENERIC FUNCTION to sum int/float values 
func SumOfNumbers[K comparable, V int64 | float64] (m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

// Declare a type constraint
type Number interface {
	int64 | float64
}

// SumNumbers is a GENERIC FUNCTION to sum int/float values using type constraint
func SumNumbers[K comparable, V Number] (m map[K]V) V {
	var s V

	for _, v := range m {
		s += v
	}

	return s
}

func main() {
	menuInts := map[string]int64{
		"eggs":    1,
		"bacon":   3,
		"sausage": 1,
	}

	menuFloats := map[string]float64{
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	fmt.Printf("[NON-GENERIC FUNCTION] Sum of Int64's %v \n", SumInts(menuInts))
	fmt.Printf("[NON-GENERIC FUNCTION] Sum of Float64's %v \n", SumFloats(menuFloats))
	fmt.Printf("[GENERIC FUNCTION] Sum of int's %v && floats %v \n",
														SumOfNumbers[string, int64](menuInts),
														SumOfNumbers[string, float64](menuFloats))

	// Remove type arguments when calling the generic function
	// Note that this isn’t always possible. For example, if you needed to call a generic function that had no arguments, you would need to include the type arguments in the function call.
	fmt.Printf("[GENERIC FUNCTION] Sum of int's %v && floats %v \n",SumOfNumbers(menuInts),SumOfNumbers(menuFloats))
	fmt.Printf("[GENERIC FUNCTION] Sum of int's %v && floats %v \n",SumNumbers(menuInts),SumNumbers(menuFloats))

}

// REFERENCE : https://go.dev/doc/tutorial/generics