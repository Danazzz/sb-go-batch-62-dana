package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// soal 1
func printMessage(kalimat string, tahun int) {
	fmt.Printf("%s %d\n", kalimat, tahun)
}

// soal 2
func kelilingSegitigaSamaSisi(sisi int, isText bool) string {
	if sisi == 0 {
		err := errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		if !isText {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic:", r)
				}
			}()
			panic(err)
		}
		return err.Error()
	}

	keliling := sisi * 3
	if isText {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling)
	}
	return fmt.Sprintf("%d", keliling)
}

// soal 3
func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}

func cetakAngka(angka *int) {
	fmt.Println("Total angka:", *angka)
}

// soal 4
func tambahPhones(phones *[]string) {
	*phones = append(*phones, "Xiaomi", "Asus", "IPhone", "Samsung", "Oppo", "Realme", "Vivo")
	sort.Strings(*phones)
	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

// soal 5
func tampilkanPhones(phones []string, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Strings(phones)
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

// soal 6
func getMovies(ch chan string, movies ...string) {
	for _, movie := range movies {
		ch <- movie
	}
	close(ch)
}

func main() {
	// soal 1
	defer printMessage("Golang Backend Development", 2021)

	// soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// soal 3
	angka := 1
	defer cetakAngka(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// soal 4
	var phones []string
	tambahPhones(&phones)

	// soal 5
	var phones2 = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup
	wg.Add(1)
	go tampilkanPhones(phones2, &wg)
	wg.Wait()

	// soal 6
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)
	go getMovies(moviesChannel, movies...)
	for value := range moviesChannel {
		fmt.Println(value)
	}
}