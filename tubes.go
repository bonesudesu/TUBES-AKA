package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// Struktur Buku
type Buku struct {
	ISBN      string
	Judul     string
	Pengarang string
}

// Fungsi Binary Search (Iteratif)
func binarySearchIterative(data []Buku, targetJudul string) (*Buku, time.Duration) {
	start := time.Now()
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2

		// Perbandingan menggunakan huruf kecil
		midJudul := strings.ToLower(data[mid].Judul)

		if midJudul == targetJudul {
			return &data[mid], time.Since(start)
		} else if midJudul < targetJudul {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return nil, time.Since(start)
}

// Fungsi Binary Search (Rekursif)
func binarySearchRecursive(data []Buku, targetJudul string, low, high int, start time.Time) (*Buku, time.Duration) {
	if low > high {
		return nil, time.Since(start)
	}

	mid := (low + high) / 2
	midJudul := strings.ToLower(data[mid].Judul)

	if midJudul == targetJudul {
		return &data[mid], time.Since(start)
	} else if midJudul < targetJudul {
		return binarySearchRecursive(data, targetJudul, mid+1, high, start)
	} else {
		return binarySearchRecursive(data, targetJudul, low, mid-1, start)
	}
}

func main() {
	// Dataset Buku
	buku := []Buku{
		{"9780007117116", "The Hobbit", "J.R.R. Tolkien"},
		{"9780131103627", "The C Programming Language", "Kernighan & Ritchie"},
		{"9780132350884", "Clean Code", "Robert C. Martin"},
		{"9780134685991", "Effective Java", "Joshua Bloch"},
		{"9780135957059", "Artificial Intelligence: A Modern Approach", "Stuart Russell"},
		{"9780201633610", "Design Patterns", "Erich Gamma"},
		{"9780321751041", "Introduction to Algorithms", "Thomas H. Cormen"},
		{"9780321982384", "Computer Networking: A Top-Down Approach", "Kurose & Ross"},
		{"9780323913037", "Deep Learning", "Ian Goodfellow"},
		{"9780596007126", "Head First Design Patterns", "Eric Freeman"},
		{"9780596158101", "JavaScript: The Good Parts", "Douglas Crockford"},
		{"9781449355739", "Learning Python", "Mark Lutz"},
		{"9781491904244", "Fluent Python", "Luciano Ramalho"},
		{"9781491954249", "Python for Data Analysis", "Wes McKinney"},
		{"9781492078005", "Data Science from Scratch", "Joel Grus"},
		{"9781617294136", "Grokking Algorithms", "Aditya Bhargava"},
		{"9781680502398", "Programming Elixir", "Dave Thomas"},
		{"9781718500457", "Rust Programming by Example", "Guillaume Gomez"},
		{"9781786466223", "Mastering Go", "Mihalis Tsoukalos"},
		{"9781789134667", "Go in Practice", "Matt Butcher"},
	}

	// Urutkan dataset secara case-insensitive berdasarkan Judul
	sort.Slice(buku, func(i, j int) bool {
		return strings.ToLower(buku[i].Judul) < strings.ToLower(buku[j].Judul)
	})

	// Input nama buku dari pengguna
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama buku yang ingin dicari: ")
	judulCari, _ := reader.ReadString('\n')
	judulCari = strings.TrimSpace(strings.ToLower(judulCari))

	// Pilihan metode pencarian
	fmt.Print("Pilih metode pencarian (1.iteratif/2.rekursif): ")
	metode, _ := reader.ReadString('\n')
	metode = strings.TrimSpace(metode)

	// Pencarian buku
	var hasil *Buku
	var durasi time.Duration
	if metode == "1" {
		hasil, durasi = binarySearchIterative(buku, judulCari)
	} else if metode == "2" {
		hasil, durasi = binarySearchRecursive(buku, judulCari, 0, len(buku)-1, time.Now())
	} else {
		fmt.Println("Metode tidak valid. Pilih iteratif atau rekursif.")
		return
	}

	// Hasil Pencarian
	if hasil != nil {
		fmt.Printf("Buku ditemukan: %s oleh %s (ISBN: %s)\n", hasil.Judul, hasil.Pengarang, hasil.ISBN)
	} else {
		fmt.Println("Buku tidak ditemukan.")
	}
	fmt.Printf("Waktu pencarian: %.10f detik\n", durasi.Seconds())
}
