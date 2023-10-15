package main

import (
	"fmt"
	"math"
)

// function soal 1
func hitungLuasandKeliling(luas *float64, keliling *float64, jariJari float64) {
	*luas = math.Pi * jariJari * jariJari
	*keliling = math.Pi * 2 * jariJari
}

// function soal 2
func introduce(sentence *string, name string, gender string, job string, age string) {
	if gender == "laki-laki" {
		*sentence = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else if gender == "perempuan" {
		*sentence = "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}

}

// function soal 3
func nambahBuah(fruits *[]string) {
	*fruits = append(*fruits, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

}

// function soal 4
func tambahDataFilm(movie string, hours string, genre string, years string, dataFilm *[]map[string]string) {
	*dataFilm = append(*dataFilm, map[string]string{
		"movie": movie,
		"hours": hours,
		"genre": genre,
		"years": years,
	})
}

func main() {

	// soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64
	var jarijariLigkaran float64

	fmt.Print("Masukkan jari-jari Lingkaran: ")
	fmt.Scan(&jarijariLigkaran)

	hitungLuasandKeliling(&luasLigkaran, &kelilingLingkaran, jarijariLigkaran)

	fmt.Println("Luas Lingkaran: ", luasLigkaran)
	fmt.Println("Keliling Lingkaran: ", kelilingLingkaran)

	// soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{}
	nambahBuah(&buah)

	for i, fruit := range buah {
		fmt.Printf("%d. %s\n", i+1, fruit)
	}

	// soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, pelem := range dataFilm {
		fmt.Printf("%d title: %s\n  duration: %s\n  genre: %s\n  year: %s\n\n", i+1, pelem["movie"], pelem["hours"], pelem["genre"], pelem["years"])
	}
}
