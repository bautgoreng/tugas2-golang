package main

import "fmt"

func isPrime(n int) bool {

	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true

}

func main() {

	var num int
	fmt.Println("Masukkan bilangan: ")
	fmt.Scanln(&num)

	if isPrime(num) {
		fmt.Println(num, "adalah bilangan prima.")
	} else {
		fmt.Println(num, "bukan bilangan prima.")
	}

}
