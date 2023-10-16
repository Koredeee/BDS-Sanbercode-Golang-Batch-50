package main

import (
	"fmt"
	// "strconv"
	// "strings"
	"sort"
	"time"
)

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

// nested struct
type student2 struct {
	person struct {
		name string
		age  int
	}
	grade int
}

// method
func (s student2) sayHello() {
	fmt.Printf("Hellom i'm %s, i am %d years old and i'm %dnd grade\n", s.person.name, s.person.age, s.grade)
}

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

	// var numberA int = 4
	// var numberB *int = &numberA
	// fmt.Println(numberB)

	// struct
	// contoh 1
	// var john = student{}
	// john.name = "john"
	// john.age = 21
	// john.grade = 2

	// fmt.Println("name  :", john.name)
	// fmt.Println("age   :", john.age)
	// fmt.Println("age   :", john.person.age)
	// fmt.Println("grade :", john.grade)

	// // contoh 2
	// var doeData = person{
	// 	name: "doe",
	// 	age:  21,
	// }
	// var doe = student{
	// 	person: doeData,
	// 	grade:  2,
	// }

	// fmt.Println("name  :", doe.name)
	// fmt.Println("age   :", doe.age)
	// fmt.Println("grade :", doe.grade)

	// anonymous struct tanpa pengisian property
	// var john = struct {
	// 	person
	// 	grade int
	// }{}
	// john.person = person{"wick", 21}
	// john.grade = 2

	// // anonymous struct dengan pengisian property
	// var doe = struct {
	// 	person
	// 	grade int
	// }{
	// 	person: person{"dor", 21},
	// 	grade:  2,
	// }

	// fmt.Println("name  :", john.person.name)
	// fmt.Println("name   :", doe.person.name)

	var johnData = person{
		name: "John",
		age:  21,
	}

	var john = student2{
		person: johnData,
		grade:  2,
	}

	john.sayHello()

	fmt.Println(time.Now())

	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)

}
