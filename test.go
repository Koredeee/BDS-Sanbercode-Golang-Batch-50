package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func introduction(firstName string, lastName string) (string, string) {
	introFirstName := "Hello My First Name Is " + firstName
	introLastName := "Hello My Last Name Is " + lastName
	return introFirstName, introLastName
}

// contoh 1
func tambahAngka(firstNumber int, lastNumber int) (jumlah int) {
	jumlah = firstNumber + lastNumber
	return
}

// contoh 2
func tampilkanKataDanAngka() (firstWord, secondWord string, number int) {
	firstWord = "Halo"
	secondWord = "Dunia"
	number = 10
	return
}

// variadic function
func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func luasSegitiga(alas int, tinggi int) int {
	return alas * tinggi / 2
}

func main() {
	// contoh 1
	// text1 := "ini teks 1"
	// text1 = "ini teks 1 diubah"
	// fmt.Println(text1)

	// // contoh 1
	// text2 := "ini teks 2"
	// fmt.Println(text2)

	// strconv.Atoi
	// var number = "124"
	// num, pante := strconv.Atoi(number)

	// if pante == nil {
	// 	fmt.Println(num)
	// } else {
	// 	fmt.Println(pante)
	// }

	// var num = 124
	// str := strconv.Itoa(num)

	// fmt.Println(str)
	// }

	// // function return multiple value
	// firstName, lastName := introduction("John", "Doe")
	// fmt.Println(firstName, lastName)

	// // bisa juga salah satunya saja meskipun func nya itu sendiri return multiple value
	// firstName2, _ := introduction("John", "Doe")
	// fmt.Println(firstName2)

	// hasil := tambahAngka(4, 5)
	// fmt.Println(hasil)

	// kataPertama, kataKedua, angka := tampilkanKataDanAngka()
	// fmt.Println(kataPertama, kataKedua, angka)

	// // variadic function
	// var total = sum(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	// fmt.Println(total)

	// variadic function with data slice
	// var numbers = []int{2, 6, 7, 8, 9, 10}
	// var total = sum(numbers...)
	// fmt.Println(total)

	// var str = "Golang"

	// for _, char := range str {
	// 	fmt.Println(strconv.Itoa(int(char)))
	// }x

	// const (
	// 	title = "pante"
	// )
	// fmt.Println(title)

	// fmt.Println(luasSegitiga(5, 6))

	// var number *int = 4

	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println(numberB)

}
