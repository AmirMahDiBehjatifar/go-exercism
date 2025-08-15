package hamming
import (
    "fmt"
)
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1 , fmt.Errorf("Hamming Length must be equal")
	} else {
        var diff int
        for i := 0 ; i < len(a) ; i++ {
            if a[i] != b[i] {
                diff++
            }
        }
        return diff , nil 
	}
}

