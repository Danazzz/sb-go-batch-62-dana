package main

import (
	"fmt"
	"strings"
)

// soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return 3.14 * t.jariJari * t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * 3.14 * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * float64((b.panjang*b.lebar)+(b.panjang*b.tinggi)+(b.lebar*b.tinggi))
}

// soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

type tampilPhone interface {
	display()
}

func (p phone) display() {
	fmt.Printf("Name: %s\nBrand: %s\nYear: %d\nColors: %s\n", p.name, p.brand, p.year, strings.Join(p.colors, ", "))
}

// soal 3
func luasPersegi(sisi int, detailed bool) interface{} {
	if sisi == 0 && detailed {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && !detailed {
		return nil
	}

	luas := sisi * sisi
	if detailed {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}

func main() {
	// soal 1
	segitiga := segitigaSamaSisi{alas: 6, tinggi: 5}
	persegi := persegiPanjang{panjang: 8, lebar: 4}
	tabung := tabung{jariJari: 7, tinggi: 10}
	balok := balok{panjang: 10, lebar: 6, tinggi: 4}

	fmt.Println("\nBangun Datar:")
	fmt.Printf("Segitiga: Luas = %d, Keliling = %d\n", segitiga.luas(), segitiga.keliling())
	fmt.Printf("Persegi Panjang: Luas = %d, Keliling = %d\n", persegi.luas(), persegi.keliling())

	fmt.Println("Bangun Ruang:")
	fmt.Printf("Tabung: Volume = %.2f, Luas Permukaan = %.2f\n", tabung.volume(), tabung.luasPermukaan())
	fmt.Printf("Balok: Volume = %.2f, Luas Permukaan = %.2f\n", balok.volume(), balok.luasPermukaan())

	// soal 2
	phone := phone{
		name:  "Galaxy S22",
		brand: "Samsung",
		year:  2022,
		colors: []string{
			"Black",
			"White",
			"Green",
		},
	}

	fmt.Println("\nPhone Info:")
	phone.display()

	// soal 3
	fmt.Println("\nLuas Persegi:")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angka1 := kumpulanAngkaPertama.([]int)
	angka2 := kumpulanAngkaKedua.([]int)
	total := 0
	var strAngka []string

	for _, num := range angka1 {
		strAngka = append(strAngka, fmt.Sprintf("%d", num))
		total += num
	}
	for _, num := range angka2 {
		strAngka = append(strAngka, fmt.Sprintf("%d", num))
		total += num
	}

	fmt.Printf("\n%s %s = %d\n", prefix.(string), strings.Join(strAngka, " + "), total)
}