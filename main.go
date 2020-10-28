package main

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}

func main() {
	for i := 0; i < 50; i++ {
		a := fibo(i)
		println(a)
	}
}
