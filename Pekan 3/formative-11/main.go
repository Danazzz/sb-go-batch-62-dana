package main

import (
	"encoding/json"
	"fmt"
)

// soal 1
func soal1() {
	var jsonString = `[{"vendor": "samsung", "model": "Galaxy S21"}, {"vendor": "sony", "model": "Xperia XZ"}, {"vendor": "samsung", "model": "Galaxy Note"}]`
	var data []map[string]string

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Data dari vendor Samsung:")
	for _, item := range data {
		if item["vendor"] == "samsung" {
			fmt.Println(item)
		}
	}
}

// soal 2
type Device struct {
	Vendor string `json:"vendor"`
	Model  string `json:"model"`
}

func soal2() {
	var jsonString = `[{"vendor": "samsung", "model": "Galaxy S21"}, {"vendor": "sony", "model": "Xperia XZ"}, {"vendor": "sony", "model": "Xperia 5"}]`
	var data []Device

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Data dari vendor Sony:")
	for _, item := range data {
		if item.Vendor == "sony" {
			fmt.Println(item)
		}
	}
}

// soal 3
type Book struct {
	Title       string `json:"title"`
	Desc        string `json:"desc"`
	Author      string `json:"author"`
	ReleaseYear int    `json:"releaseYear"`
}

func soal3() {
	books := []Book{
		{"Go Programming", "Learn Go from scratch", "John Doe", 2021},
		{"Mastering Python", "Advanced Python techniques", "Jane Smith", 2020},
		{"Java Basics", "Introduction to Java", "James Brown", 2019},
		{"Web Development", "Full-stack web development", "Michael Johnson", 2022},
		{"Data Science", "Machine learning and AI", "Emily Davis", 2023},
	}

	jsonData, err := json.Marshal(books)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Books in JSON format:")
	fmt.Println(string(jsonData))
}

func main() {
	fmt.Println("Soal 1:")
	soal1()

	fmt.Println("\nSoal 2:")
	soal2()

	fmt.Println("\nSoal 3:")
	soal3()
}