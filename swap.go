package problems

func Swap(num1 int, num2 int) (int, int) {
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2

	return num1, num2
}
