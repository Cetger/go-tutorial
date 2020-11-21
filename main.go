package main

import (
	"fmt"
	"os"
	"strconv"
)

//https://tour.golang.org/flowcontrol/8 solution of this question
func Sqrt(x float64) float64 {
	z := 1.0
	_ = z
	for counter := 1; counter < 10; counter++ {
		fmt.Println("X : ", x, "Z : ", z, " total: ", x-z*z)
		if x-z*z < 0.001 && x-z*z >= 0 {
			fmt.Println("Eureka: ", z)
			return z
		} else {
			z -= (z*z - x) / (2 * z)
		}
	}
	return z
}

func main() {
	if s, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		fmt.Printf("Result : %v\n", Sqrt(s))
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
