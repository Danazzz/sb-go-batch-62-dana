package main

import "fmt"

// Soal 1
func hitungLingkaran(luas *float64, keliling *float64, jariJari float64) {
	const pi = 3.14
	*luas = pi * jariJari * jariJari
	*keliling = 2 * pi * jariJari
}

// Soal 2
func introduce(sentence *string, nama, gender, pekerjaan, umur string) {
	var sapaan string
	if gender == "laki-laki" {
		sapaan = "Pak"
	} else if gender == "perempuan" {
		sapaan = "Bu"
	}
	*sentence = fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", sapaan, nama, pekerjaan, umur)
}

// Soal 3
func tambahBuah(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}

// Soal 4
func tambahDataFilm(judul, durasi, genre, tahun string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":    judul,
		"duration": durasi,
		"genre":    genre,
		"year":     tahun,
	}
	*dataFilm = append(*dataFilm, film)
}

func main() {
	// Soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64
	hitungLingkaran(&luasLingkaran, &kelilingLingkaran, 9)
	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	// Soal 2
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)
	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// Soal 3
	var buah = []string{}
	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	for i, b := range buah {
		fmt.Printf("%d. %s\n", i+1, b)
	}

	// Soal 4
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. ", i+1)
		for key, value := range film {
			fmt.Printf("%s : %s \n", key, value)
		}
		fmt.Println()
	}
}