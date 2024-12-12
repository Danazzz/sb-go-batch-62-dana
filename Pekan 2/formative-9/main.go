package main

import (
	"fmt"
	"strings"
	"formative9/formative8"
)

func main() {
	// soal 1
	segitiga := formative8.SegitigaSamaSisi{Alas: 6, Tinggi: 5}
	persegi := formative8.PersegiPanjang{Panjang: 8, Lebar: 4}
	tabung := formative8.Tabung{JariJari: 7, Tinggi: 10}
	balok := formative8.Balok{Panjang: 10, Lebar: 6, Tinggi: 4}

	fmt.Println("\nBangun Datar:")
	fmt.Printf("Segitiga: Luas = %d, Keliling = %d\n", segitiga.Luas(), segitiga.Keliling())
	fmt.Printf("Persegi Panjang: Luas = %d, Keliling = %d\n", persegi.Luas(), persegi.Keliling())

	fmt.Println("Bangun Ruang:")
	fmt.Printf("Tabung: Volume = %.2f, Luas Permukaan = %.2f\n", tabung.Volume(), tabung.LuasPermukaan())
	fmt.Printf("Balok: Volume = %.2f, Luas Permukaan = %.2f\n", balok.Volume(), balok.LuasPermukaan())

	// soal 2
	phone := formative8.Phone{
		Name:  "Galaxy S22",
		Brand: "Samsung",
		Year:  2022,
		Colors: []string{
			"Black",
			"White",
			"Green",
		},
	}

	fmt.Println("\nPhone Info:")
	phone.Display()

	// soal 3
	fmt.Println("\nLuas Persegi:")
	fmt.Println(formative8.LuasPersegi(4, true))
	fmt.Println(formative8.LuasPersegi(8, false))
	fmt.Println(formative8.LuasPersegi(0, true))
	fmt.Println(formative8.LuasPersegi(0, false))

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