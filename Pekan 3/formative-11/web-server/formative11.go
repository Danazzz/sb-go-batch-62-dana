package main

import (
	"fmt"
	"math"
	"net/http"
)

func calculateVolume(radius, height float64) float64 {
	return math.Pi * math.Pow(radius, 2) * height
}

func calculateBaseArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func calculateCircumference(radius float64) float64 {
	return 2 * math.Pi * radius
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		radius := 7.0
		height := 10.0

		volume := calculateVolume(radius, height)
		baseArea := calculateBaseArea(radius)
		circumference := calculateCircumference(radius)

		response := fmt.Sprintf("jariJari : %.0f, tinggi: %.0f, volume : %.2f, luas alas: %.2f, keliling alas: %.2f",
			radius, height, volume, baseArea, circumference)

		fmt.Fprintln(w, response)
	})

	fmt.Println("Web server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}