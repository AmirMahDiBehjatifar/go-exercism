package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(Valid("6037997553008906"))
}

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	validNumberRegex := regexp.MustCompile(`^[0-9]{2,}$`)

	if !validNumberRegex.MatchString(id) {
		return false
	} else {
		id = strings.ReplaceAll(id, " ", "")
		sum := 0
		doubleNext := false

		for i := len(id) - 1; i >= 0; i-- {
			digit := int(id[i] - '0')
			if doubleNext {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
			doubleNext = !doubleNext
		}
		if sum%10 == 0 {
			return true
		} else {
			return false
		}
	}
}

// strings of length 1 or less are not valid
// remove spaces
// all other non-digit characters are disallowed

// step 1 double every second digit from right
// if doubling the number results in a number greater than 9  then (subtract -9)

// then sum all of the digits

// if the sum is evenly divisible by 10 -- then number is valid
