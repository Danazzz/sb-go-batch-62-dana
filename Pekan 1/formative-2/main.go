package main

import (
	"fmt"
	"strconv"
	"strings"
	"golang.org/x/text/cases"
    "golang.org/x/text/language"
)

func main() {
	// soal 1
	var kata1 = "Bootcamp"
	var kata2 = "Digital"
	var kata3 = "Skill"
	var kata4 = "Sanbercode"
	var kata5 = "Golang"
	fmt.Println(kata1 + " " + kata2 + " " + kata3 + " " + kata4 + " " + kata5)

	// soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println(halo)

	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"
	c := cases.Title(language.Und)
	output := fmt.Sprintf("%s %s %s %s", kataPertama, c.String(kataKedua), kataKetiga[:len(kataKetiga)-1]+"R", strings.ToUpper(kataKeempat))
	fmt.Println(output)

	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"
	num1, _ := strconv.Atoi(angkaPertama)
	num2, _ := strconv.Atoi(angkaKedua)
	num3, _ := strconv.Atoi(angkaKetiga)
	num4, _ := strconv.Atoi(angkaKeempat)
	sum := num1 + num2 + num3 + num4
	fmt.Println(sum)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021
	kalimat = strings.Replace(kalimat, "halo", "Hi", 2)
	fmt.Printf("%s - %d\n", kalimat, angka)
}