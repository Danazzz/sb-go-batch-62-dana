package main

import (
	"fmt"
)

// Soal 1
type buah struct {
	nama      string
	warna     string
	adaBijinya bool
	harga     int
}

// Soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() float64 {
	return 0.5 * float64(s.alas) * float64(s.tinggi)
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

// Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(color string) {
	p.colors = append(p.colors, color)
}

// Soal 4
type movie struct {
	title, genre string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	newMovie := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}
	*dataFilm = append(*dataFilm, newMovie)
}

func main() {
	// Soal 1
	buahList := []buah{
		{"Nanas", "Kuning", false, 9000},
		{"Jeruk", "Oranye", true, 8000},
		{"Semangka", "Hijau & Merah", true, 10000},
		{"Pisang", "Kuning", false, 5000},
	}
	for _, b := range buahList {
		ket := "Ada"
		if !b.adaBijinya {
			ket = "Tidak"
		}
		fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %v, Harga: %d\n", b.nama, b.warna, ket, b.harga)
	}

	// Soal 2
	luasSegitiga := segitiga{alas: 10, tinggi: 5}
	luasPersegi := persegi{sisi: 4}
	luasPersegiPanjang := persegiPanjang{panjang: 8, lebar: 3}

	fmt.Printf("Luas Segitiga: %.2f\n", luasSegitiga.luas())
	fmt.Printf("Luas Persegi: %d\n", luasPersegi.luas())
	fmt.Printf("Luas Persegi Panjang: %d\n", luasPersegiPanjang.luas())

	// Soal 3
	myPhone := phone{name: "iPhone", brand: "Apple", year: 2022, colors: []string{"Black"}}
	myPhone.addColor("White")
	myPhone.addColor("Blue")
	fmt.Printf("Phone: %s\nBrand: %s\nYear: %d\nColors: %v\n", myPhone.name, myPhone.brand, myPhone.year, myPhone.colors)

	// Soal 4

	var dataFilm = []movie{}
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title : %s\nduration : %d jam\ngenre : %s\nyear : %d\n\n", i+1, film.title, film.duration/60, film.genre, film.year)
	}
}