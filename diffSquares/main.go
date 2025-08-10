package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SquareOfSum(6))  // expected 441
	fmt.Println(SumOfSquares(6)) // expected 91
	fmt.Println(Difference(6))   // 350
}

func SquareOfSum(n int) int {
	// ۱‌+۲+۳ = 6 -> 6*6 = 36
	var sum int = 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	result := math.Pow(float64(sum), 2)
	return int(result)
}
func SumOfSquares(n int) int {
	// 1*1 + 2*2 + 3*3 = 14
	var sum float64 = 0
	for i := 0; i <= n; i++ {
		sum += math.Pow(float64(i), 2)
	}
	return int(sum)

}
func Difference(n int) int {
	// 6 => 441 - 91 = 350
	return SquareOfSum(n) - SumOfSquares(n)
}
