package main

func fibo(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibo(n-1) + fibo(n-2)
}

func fiboPara(n int, done chan int) {
	cur := fibo(n)
	done <- cur
}

func main() {

	for i := 0; i < 45; i++ {
		a := fibo(i)
		println(a)
	}

	/*
		done := make(chan int, 45)

		for i := 0; i < 45; i++ {
			go fiboPara(i, done)
		}

		for i := 0; i < 45; i++ {
			res := <-done
			println(res)
		}
	*/
}
