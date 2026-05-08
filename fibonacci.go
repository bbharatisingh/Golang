package problems

func Fibonacci(n int) int {

	if n <= 1 {
		return n
	}
	res := Fibonacci((n - 1)) + Fibonacci(n-2)
	// fmt.Println("fib num of n:  ", res)
	return res
}
