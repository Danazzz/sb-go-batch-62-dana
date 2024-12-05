package main

import "fmt"

func main() {
    // soal 1
    fmt.Println("Soal 1:")
    for i := 1; i <= 20; i++ {
        if i%2 == 1 && i%3 == 0 {
            fmt.Printf("%d - I Love Coding\n", i)
        } else if i%2 == 0 {
            fmt.Printf("%d - Berkualitas\n", i)
        } else {
            fmt.Printf("%d - Santai\n", i)
        }
    }

    // soal 2
    fmt.Println("\nSoal 2:")
    for i := 1; i <= 7; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }

    // soal 3
    fmt.Println("\nSoal 3:")
    var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
    var result = kalimat[2:]
    fmt.Println(result)

    // soal 4
    fmt.Println("\nSoal 4:")
    var sayuran = []string{}
    sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

    for i, s := range sayuran {
        fmt.Printf("%d. %s\n", i+1, s)
    }

    // soal 5
    fmt.Println("\nSoal 5:")
    var satuan = map[string]int{
        "panjang": 7,
        "lebar":   4,
        "tinggi":  6,
    }

    volume := 1
    for key, value := range satuan {
        fmt.Printf("%s = %d\n", key, value)
        volume *= value
    }
    fmt.Printf("volume balok = %d\n", volume)
}