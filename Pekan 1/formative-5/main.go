package main

import "fmt"

func main() {
	// soal 1
	luasPersegiPanjang := func(panjang, lebar int) int {
		return panjang * lebar
	}

	kelilingPersegiPanjang := func(panjang, lebar int) int {
		return 2 * (panjang + lebar)
	}

	volumeBalok := func(panjang, lebar, tinggi int) int {
		return panjang * lebar * tinggi
	}

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
	introduce := func(nama, gender, pekerjaan, umur string) string {
		if gender == "laki-laki" {
			return fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", nama, pekerjaan, umur)
		} else if gender == "perempuan" {
			return fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", nama, pekerjaan, umur)
		} else {
			return fmt.Sprintf("Gender tidak dikenal")
		}
	}

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// soal 3
	buahFavorit := func(nama string, buah ...string) string {
		result := "Halo nama saya " + nama + " dan buah favorit saya adalah \""
		for i, val := range buah {
			if i > 0 {
				result += "\", \""
			}
			result += val
		}
		result += "\""
		return result
	}

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	// soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm := func(title, jam, genre, tahun string) {
		dataFilm = append(dataFilm, map[string]string{
			"title": title,
			"jam":   jam,
			"genre": genre,
			"tahun": tahun,
		})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}