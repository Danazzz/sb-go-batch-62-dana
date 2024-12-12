package formative8

import (
	"fmt"
	"strings"
)

// soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return 3.14 * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * float64((b.Panjang*b.Lebar)+(b.Panjang*b.Tinggi)+(b.Lebar*b.Tinggi))
}

// soal 2
type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type TampilPhone interface {
	Display()
}

func (p Phone) Display() {
	fmt.Printf("Name: %s\nBrand: %s\nYear: %d\nColors: %s\n", p.Name, p.Brand, p.Year, strings.Join(p.Colors, ", "))
}

// soal 3
func LuasPersegi(sisi int, detailed bool) interface{} {
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