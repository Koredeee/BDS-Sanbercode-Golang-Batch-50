package main

import (
	"fmt"
)

// function soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// function soal 2
func introduce(name string, gender string, job string, age string) (introducing string) {
	if gender == "laki-laki" {
		introducing = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else if gender == "perempuan" {
		introducing = "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}

	return introducing

}

// function soal 3
func buahFavorit(name string, fruits ...string) string {
	var buahFavorit = ""
	for _, fruit := range fruits {
		buahFavorit += "\"" + fruit + "\"" + ", "
	}

	return "halo nama saya " + name + " dan buah favorit saya adalah " + buahFavorit
}

// function soal 4

func main() {

	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(movie string, hours string, genre string, years string) {
		isiDataFilm := map[string]string{
			"genre": genre,
			"jam":   hours,
			"tahun": years,
			"title": movie,
		}
		dataFilm = append(dataFilm, isiDataFilm)

	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
