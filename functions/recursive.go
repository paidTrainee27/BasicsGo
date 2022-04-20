package main

func getFactorial(num int) int  {
	if num < 0 {
		return 0
	}
	if num == 0 || num == 1 {
		return 1
	}
	return num * getFactorial(num - 1)
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