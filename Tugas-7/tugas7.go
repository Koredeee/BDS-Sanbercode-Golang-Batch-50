package main

import (
	"fmt"
)

// struct soal 1
type buah struct {
	nama  string
	warna string
	biji  bool
	harga int
}

// struct soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

// method soal 2
func (s segitiga) luasSegitiga() int {
	return s.alas * s.tinggi / 2
}

func (p persegi) luasPersegi() int {
	return p.sisi * p.sisi
}

func (p persegiPanjang) luasPersegiPanjnag() int {
	return p.panjang * p.lebar
}

// struct soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addPhoneColor(color ...string) {
	p.colors = append(p.colors, color...)
}

func (p *phone) printPhoneData() {
	fmt.Println("Name\t:", p.name)
	fmt.Println("Brand\t:", p.brand)
	fmt.Println("Year\t:", p.year)
	fmt.Println("Colors\t:", p.colors)
}

// struct soal 4
type movie struct {
	title, genre    string
	duration, years int
}

func tambahDataFilm(titleMovie string, durationMovie int, genreMovie string, yearsMovie int, dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{
		title:    titleMovie,
		duration: durationMovie,
		genre:    genreMovie,
		years:    yearsMovie,
	})
}

func main() {

	// soal 1
	var nanas = buah{
		nama:  "Nanas",
		warna: "Kuning",
		biji:  false,
		harga: 9000,
	}

	var jeruk = buah{
		nama:  "Jeruk",
		warna: "Oranye",
		biji:  true,
		harga: 8000,
	}

	var semangka = buah{
		nama:  "Semangka",
		warna: "Hijau & Merah",
		biji:  true,
		harga: 10000,
	}

	var pisang = buah{
		nama:  "Pisang",
		warna: "Kuning",
		biji:  false,
		harga: 5000,
	}

	fmt.Println(nanas)
	fmt.Println(jeruk)
	fmt.Println(semangka)
	fmt.Println(pisang)

	// soal 2
	var segitigaSiji = segitiga{
		alas:   10,
		tinggi: 5,
	}

	var persegiSiji = persegi{
		sisi: 10,
	}

	var persegiPanjangSiji = persegiPanjang{
		panjang: 10,
		lebar:   5,
	}

	fmt.Println(persegiPanjangSiji.luasPersegiPanjnag())
	fmt.Println(segitigaSiji.luasSegitiga())
	fmt.Println(persegiSiji.luasPersegi())

	// soal 3
	var sumsang = phone{
		name:   "Sumsang Galaksi S100",
		brand:  "Sumsang",
		year:   2020,
		colors: []string{"Pink", "Red", "Yellow", "Green"},
	}

	sumsang.addPhoneColor("Black", "White", "Blue")
	sumsang.printPhoneData()

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d title: %s\n  duration: %d\n  genre: %s\n  year: %d\n\n", i+1, film.title, film.duration, film.genre, film.years)
	}

}
