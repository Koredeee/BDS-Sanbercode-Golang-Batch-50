package main

import (
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// converting
	var panjangPersegiPanjangConv, _ = strconv.Atoi(panjangPersegiPanjang)
	var lebarPersegiPanjangConv, _ = strconv.Atoi(lebarPersegiPanjang)

	var alasSegitigaConv, _ = strconv.Atoi(alasSegitiga)
	var tinggiSegitigaConv, _ = strconv.Atoi(tinggiSegitiga)

	// answers
	var luasPersegiPanjang int = panjangPersegiPanjangConv * lebarPersegiPanjangConv
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangConv + lebarPersegiPanjangConv)
	var luasSegitiga int = (alasSegitigaConv * tinggiSegitigaConv) / 2

	fmt.Printf("Luas Persegi Panjang: %d\n", luasPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang: %d\n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga: %d\n", luasSegitiga)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	// index nilai John
	var idxJohn string
	if nilaiJohn >= 90 {
		idxJohn = "A"
		fmt.Println(idxJohn)
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		idxJohn = "B"
		fmt.Println(idxJohn)
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		idxJohn = "C"
		fmt.Println(idxJohn)
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		idxJohn = "D"
		fmt.Println(idxJohn)
	} else if nilaiJohn < 50 {
		idxJohn = "E"
		fmt.Println(idxJohn)
	}

	// index nilai Doe
	var idxDoe string
	if nilaiDoe >= 90 {
		idxDoe = "A"
		fmt.Println(idxDoe)
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		idxDoe = "B"
		fmt.Println(idxDoe)
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		idxDoe = "C"
		fmt.Println(idxDoe)
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		idxDoe = "D"
		fmt.Println(idxDoe)
	} else if nilaiDoe < 50 {
		idxDoe = "E"
		fmt.Println(idxDoe)
	}

	// soal 3
	var tanggal = 17
	var bulan = 8
	var tahun = 1945

	tanggal = 7
	bulan = 2
	tahun = 2003

	// converted to string
	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	// soal 4
	lahir := 2003

	var gen string

	if lahir >= 1944 && lahir <= 1964 {
		gen = "Baby boomer"
	} else if lahir >= 1965 && lahir <= 1979 {
		gen = "Generasi X"
	} else if lahir >= 1980 && lahir <= 1994 {
		gen = "Generasi Y (Millennials)"
	} else if lahir >= 1995 && lahir <= 2015 {
		gen = "Generasi Z"
	} else {
		gen = "Generasi tidak terdefinisi"
	}

	fmt.Printf("Tahun kelahiran Anda (%d) termasuk dalam gen: %s\n", lahir, gen)

}
