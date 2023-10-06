package main

import (
	"fmt"
)

func main() {
	// soal 1
	for number := 1; number <= 20; number++ {
		// cek number -> kelipatan 3 dan ganjil
		if number%3 == 0 && number%2 != 0 {
			fmt.Printf("%d - I Love Coding\n", number)
		} else if number%2 != 0 { // cek number -> ODD
			fmt.Printf("%d - Santai\n", number)
		} else if number%2 == 0 { // cek number -> EVEN
			fmt.Printf("%d - Berkualitas\n", number)
		}
	}

	// soal 2
	tinggiAlas := 7

	for i := 1; i <= tinggiAlas; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("#")
		}

		fmt.Println()
	}

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// soal 4
	var sayuran = []string{}
	var addSayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, sayuranGes := range addSayuran {
		fmt.Printf("%d. %s\n", i+1, sayuranGes)
	}

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	var volume = 1

	for _, hasil := range satuan {
		volume *= hasil
	}

	fmt.Printf("panjang = %d\n", satuan["panjang"])
	fmt.Printf("lebar = %d\n", satuan["lebar"])
	fmt.Printf("tinggi = %d\n", satuan["tinggi"])
	fmt.Printf("Volume Balok = %d\n", volume)

}
