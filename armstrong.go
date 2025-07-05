// An Armstrong number (also known as a narcissistic number) is a number that is equal to the sum of its own digits each raised to the power of the number of digits.
//
//	For example, 153 is an Armstrong number because (1^3 + 5^3 + 3^3 = 153).
package problems

import (
	"fmt"
	"math"
	"strconv"
)

func CheckArmstrong() {
	num := 153
	n := len(strconv.Itoa(num))
	sum := 0
	for num > 0 {
		sum += int(math.Pow(float64(num%10), float64(n)))
		num = num / 10
	}
	fmt.Println(sum)
}
