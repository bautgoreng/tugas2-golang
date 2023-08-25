package main

import "fmt"

func main() {

	var n int
	fmt.Println("Masukan batasan: ")
	fmt.Scanln(&n)

	fmt.Println("Bilangan kelipatan 7 antara 1 dan n:")

	for i := 1; i <= n; i++ {

		if i%7 == 0 {

			fmt.Println(i)

		}

	}

}
