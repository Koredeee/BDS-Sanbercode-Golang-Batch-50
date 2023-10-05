// Program to illustrate fmt.Print()

package main

// import the fmt package
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// soal 1
	golanggo1 := "Bootcamp"
	golanggo2 := "Digital"
	golanggo3 := "Skill"
	golanggo4 := "Sanbercode"
	golanggo5 := "Golang"

	fmt.Println(golanggo1, golanggo2, golanggo3, golanggo4, golanggo5)

	// // soal 2
	halo := "Halo Dunia"
	halo = "Halo Golang"
	fmt.Print(halo)

	// // soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var kataKeduaFind = "s"
	var kataKetigaFind = "r"

	// fmt.Println(kataPertama)

	kataKeduaMod := strings.Replace(kataKedua, kataKeduaFind, strings.ToUpper(kataKeduaFind), 1)
	// fmt.Println(kataKeduaMod)

	kataKetigaMod := strings.Replace(kataKetiga, kataKetigaFind, strings.ToUpper(kataKetigaFind), 1)
	// fmt.Println(kataKetigaMod)

	fmt.Println(kataPertama, kataKeduaMod, kataKetigaMod, strings.ToUpper(kataKeempat))

	// // soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	// converting to int
	var angkaPertamaConv, _ = strconv.Atoi(angkaPertama)
	var angkaKeduaConv, _ = strconv.Atoi(angkaKedua)
	var angkaKetigaConv, _ = strconv.Atoi(angkaKetiga)
	var angkaKeempatConv, _ = strconv.Atoi(angkaKeempat)

	fmt.Println(angkaPertamaConv + angkaKeduaConv + angkaKetigaConv + angkaKeempatConv)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimatMod := strings.Replace(kalimat, "halo", "Hi", -1)

	fmt.Printf("\"%s\" - %d", kalimatMod, angka)

}
