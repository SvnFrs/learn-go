package functionality

func Multiply(a int, b int) int {
	status := ZeroChecker(a, b)

	if status {
		return a * b
	}
	return 0
}
