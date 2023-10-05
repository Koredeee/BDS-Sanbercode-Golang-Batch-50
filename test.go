package main

import (
	"fmt"
	"strconv"
)

func main() {
	// contoh 1
	// text1 := "ini teks 1"
	// text1 = "ini teks 1 diubah"
	// fmt.Println(text1)

	// // contoh 1
	// text2 := "ini teks 2"
	// fmt.Println(text2)

	// strconv.Atoi
	var number = "124"
	num, pante := strconv.Atoi(number)

	if pante == nil {
		fmt.Println(num)
	} else {
		fmt.Println(pante)
	}

	// var num = 124
	// str := strconv.Itoa(num)

	// fmt.Println(str)
	// }
}
