package main

func getFactorial(num int) int {
	if num < 0 {
		return 0
	}
	if num == 0 || num == 1 {
		return 1
	}
	return num * getFactorial(num-1)
}

func getFib(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return getFib(n-1) + getFib(n-2)
}

func getFactorialSimple(n int) int {
	fact := 1
	if n < 0 {
		return 0
	}

	if n == 0 || n == 1 {
		return fact
	}

	for i := n; i > 1; i-- {
		fact *= i
	}

	return fact
}
