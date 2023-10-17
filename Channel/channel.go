package main

import (
	"fmt"
	"sync"
)

func cetak(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}

func printText(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
	wg.Done()
}

func main() {
	// ch := make(chan int)

	// go cetak(ch)

	// for {
	// 	data, ok := <-ch
	// 	if ok == false {
	// 		break
	// 	}
	// 	fmt.Printf("Data di terima %v\b\n", data)
	// }

	// ch := make(chan int, 5)

	// ch <- 6
	// ch <- 7
	// ch <- 5
	// ch <- 5

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	// // channel select
	// var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	// fmt.Println("numbers :", numbers)

	// var ch1 = make(chan float64)
	// go getAverage(numbers, ch1)

	// var ch2 = make(chan int)
	// go getMax(numbers, ch2)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case avg := <-ch1:
	// 		fmt.Printf("Avg \t: %.2f \n", avg)
	// 	case max := <-ch2:
	// 		fmt.Printf("Max \t: %d \n", max)
	// 	}
	// }

	// waitGroup
	var wg sync.WaitGroup

	wg.Add(1)
	go printText("Halo", &wg)

	wg.Add(1)
	go printText("Dunia", &wg)

	wg.Wait()
}
