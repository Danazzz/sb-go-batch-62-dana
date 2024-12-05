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

	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int = panjang * lebar
	var kelilingPersegiPanjang int = 2 * (panjang + lebar)
	var luasSegitiga int = (alas * tinggi) / 2

	fmt.Println("Luas Persegi Panjang:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Println("Indeks nilai John:", getIndeks(nilaiJohn))
	fmt.Println("Indeks nilai Doe:", getIndeks(nilaiDoe))

	// soal 3
	var tanggal = 18
	var bulan = 3
	var tahun = 1997

	var bulanString string
	switch bulan {
	case 1:
		bulanString = "Januari"
	case 2:
		bulanString = "Februari"
	case 3:
		bulanString = "Maret"
	case 4:
		bulanString = "April"
	case 5:
		bulanString = "Mei"
	case 6:
		bulanString = "Juni"
	case 7:
		bulanString = "Juli"
	case 8:
		bulanString = "Agustus"
	case 9:
		bulanString = "September"
	case 10:
		bulanString = "Oktober"
	case 11:
		bulanString = "November"
	case 12:
		bulanString = "Desember"
	default:
		bulanString = "Bulan tidak valid"
	}
	fmt.Printf("%d %s %d\n", tanggal, bulanString, tahun)

	// soal 4
	var tahunLahir = 1997
	if tahunLahir >= 1944 && tahunLahir <= 1964 {
		fmt.Println("Generasi: Baby Boomer")
	} else if tahunLahir >= 1965 && tahunLahir <= 1979 {
		fmt.Println("Generasi: Generasi X")
	} else if tahunLahir >= 1980 && tahunLahir <= 1994 {
		fmt.Println("Generasi: Generasi Y (Millenials)")
	} else if tahunLahir >= 1995 && tahunLahir <= 2015 {
		fmt.Println("Generasi: Generasi Z")
	} else {
		fmt.Println("Generasi tidak terdefinisi")
	}
}

func getIndeks(nilai int) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 && nilai < 80 {
		return "B"
	} else if nilai >= 60 && nilai < 70 {
		return "C"
	} else if nilai >= 50 && nilai < 60 {
		return "D"
	} else {
		return "E"
	}
}