package main

import (
	"log"
	"net/http"
	httpresp "practice_set/httpReq"
)

//"practice_set/concurrency"

// Matches the module name defined in your go.mod

// func checkPlndrome(number int) bool {
// 	// Convert the integer to a string

// 	num := strconv.Itoa(number)

// 	i, j := 0, len(num)-1
// 	for i < j {
// 		if num[i] != num[j] {
// 			return false
// 		}
// 		i += 1
// 		j -= 1
// 	}
// 	return true
// }

func main() {
	// num,i := 12364321, 0
	// fmt.Println(checkPlndrome(num))a
	// armstrong.CheckArmstrong()
	// num = 3
	// for i := 1; i < 10; i++ {
	// 	fmt.Println(problems.Factorial(i))
	// }
	// fmt.Println(problems.Swap(2, 3))
	// concurrency.Fibnacci_num([]int{1, 2, 3, 45})
	responseList := httpresp.BasicOP([]string{"https://www.example.com",
		"https://www.google.com",
		"https://www.github.com"})
	for _, resp := range responseList {
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("%v", resp.Body)
		}
	}

}
