package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

// function soal 1
func printArray(phones []string, wg *sync.WaitGroup) {
	for i, e := range phones {
		fmt.Printf("%d. %s\n", i+1, e)
	}

	sort.Strings(phones)
	wg.Done()
}

// function soal 2
func getMovies(moviesChannel chan string, movies ...string) {
	for _, e := range movies {
		moviesChannel <- e
	}
	close(moviesChannel)
}

// function soal 3
func hitungLuasLingkaran(jariJari float64, hasilLuasChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	luas := math.Pi * jariJari * jariJari
	hasilLuasChan <- luas
}

func hitungKelilingLingkaran(jariJari float64, hasilKelilingChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	keliling := 2 * math.Pi * jariJari
	hasilKelilingChan <- keliling
}

func hitungVolumeTabung(jariJari, tinggi float64, hasilVolumeChan chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	volume := math.Pi * jariJari * jariJari * tinggi
	hasilVolumeChan <- volume
}

func hitungLuasPersegiPanjang(panjang, lebar float64, hasilLuasChan chan float64) {
	luas := panjang * lebar
	hasilLuasChan <- luas
}

func hitungKelilingPersegiPanjang(panjang, lebar float64, hasilKelilingChan chan float64) {
	keliling := 2 * (panjang + lebar)
	hasilKelilingChan <- keliling
}

func hitungVolumeBalok(panjang, lebar, tinggi float64, hasilVolumeChan chan float64) {
	volume := panjang * lebar * tinggi
	hasilVolumeChan <- volume

}

func main() {
	// soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)
	var wg sync.WaitGroup

	wg.Add(1)
	go printArray(phones, &wg)

	wg.Wait()

	// soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	fmt.Println("List Movies:")
	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	jariJari := []float64{8, 14, 20}
	tinggiTabung := 10

	hasilLuasChan := make(chan float64, len(jariJari))
	hasilKelilingChan := make(chan float64, len(jariJari))
	hasilVolumeChan := make(chan float64, len(jariJari))

	for _, r := range jariJari {
		wg.Add(1)
		go hitungLuasLingkaran(r, hasilLuasChan, &wg)

		wg.Add(1)
		go hitungKelilingLingkaran(r, hasilKelilingChan, &wg)

		wg.Add(1)
		go hitungVolumeTabung(r, float64(tinggiTabung), hasilVolumeChan, &wg)
	}

	wg.Wait()
	close(hasilLuasChan)
	close(hasilKelilingChan)
	close(hasilVolumeChan)

	for r := range hasilLuasChan {
		fmt.Printf("Luas Lingkaran dengan jari-jari %.2f: %.2f\n", r, <-hasilLuasChan)
	}

	for r := range hasilKelilingChan {
		fmt.Printf("Keliling Lingkaran dengan jari-jari %.2f: %.2f\n", r, <-hasilKelilingChan)
	}

	for r := range hasilVolumeChan {
		fmt.Printf("Volume Tabung dengan jari-jari %.2f dan tinggi %d: %.2f\n", r, tinggiTabung, <-hasilVolumeChan)
	}

	// soal 4
	panjang := 5.0
	lebar := 3.0
	tinggi := 2.0

	hasilLuasPersegiChan := make(chan float64)
	hasilKelilingPersegiChan := make(chan float64)
	hasilVolumeBalokChan := make(chan float64)

	wg.Add(1)
	go func() {
		hitungLuasPersegiPanjang(panjang, lebar, hasilLuasPersegiChan)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		hitungKelilingPersegiPanjang(panjang, lebar, hasilKelilingPersegiChan)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		hitungVolumeBalok(panjang, lebar, tinggi, hasilVolumeBalokChan)
		wg.Done()
	}()

	go func() {
		close(hasilLuasPersegiChan)
		close(hasilKelilingPersegiChan)
		close(hasilVolumeBalokChan)
	}()
	wg.Wait()

	wg.Done()

	select {
	case luas := <-hasilLuasPersegiChan:
		fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
	case keliling := <-hasilKelilingPersegiChan:
		fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)
	case volume := <-hasilVolumeBalokChan:
		fmt.Printf("Volume Balok: %.2f\n", volume)
	}
}
