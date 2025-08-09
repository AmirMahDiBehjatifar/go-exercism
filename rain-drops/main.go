package main

import "fmt"

func main() {
	fmt.Println(Convert(105))
}
func Convert(number int) string {
	// if
	switch {
	case number%3 == 0 && number%5 == 0 && number%7 == 0:
		return "PlingPlangPlong"
	case number%3 == 0 && number%5 == 00:
		return "PlingPlang"
	case number%3 == 0 && number%7 == 0:
		return "PlingPlong"
	case number%5 == 0 && number%7 == 0:
		return "PlangPlong"
	case number%3 == 0:
		return "Pling"

	case number%5 == 0:
		return "Plang"

	case number%7 == 0:
		return "Plong"

	default:
		stringVal := fmt.Sprintf("%d", number)
		return stringVal
	}

}
