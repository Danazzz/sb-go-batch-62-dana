package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var nextID uint = 1

// Middleware
func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "admin" || password != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Username atau Password tidak sesuai"))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func hitungIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

// soal 1
func postNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		var input struct {
			Nama       string `json:"nama"`
			MataKuliah string `json:"mata_kuliah"`
			Nilai      uint   `json:"nilai"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
			return
		}

		if input.Nilai > 100 {
			http.Error(w, "Nilai maksimal adalah 100", http.StatusBadRequest)
			return
		}

		mahasiswaBaru := NilaiMahasiswa{
			ID:         nextID,
			Nama:       input.Nama,
			MataKuliah: input.MataKuliah,
			Nilai:      input.Nilai,
			IndeksNilai: hitungIndeksNilai(input.Nilai),
		}
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, mahasiswaBaru)
		nextID++

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(mahasiswaBaru)
		return
	}
	http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
}

// soal 2
func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
		return
	}
	http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/post_nilai_mahasiswa", BasicAuth(postNilaiMahasiswa))
	http.HandleFunc("/get_nilai_mahasiswa", getNilaiMahasiswa)

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}