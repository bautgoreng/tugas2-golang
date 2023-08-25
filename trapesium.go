package main

import "fmt"

func main() {
	var a float64
	var b float64
	var t float64
	var res float64

	fmt.Println("Masukkan alas: ")
	fmt.Scanln(&a)
	fmt.Println("Masukkan sisi miring: ")
	fmt.Scanln(&b)
	fmt.Println("Masukkan tinggi: ")
	fmt.Scanln(&t)

	res = 0.5 * (a + b) * t
	fmt.Println("Luas trapesium adalah: ", res)
}
