package main

import (
	"fmt"
	"strconv"
	"strings"
)

func a000124(n int) []int {
	result := []int{}
	current := 1
	for i := 0; i < n; i++ {
		result = append(result, current)
		current += i + 1
	}
	return result
}

func main() {
	var input int
	fmt.Print("Masukkan jumlah elemen: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Input tidak valid. Masukkan angka bulat.")
		return
	}

	sequence := a000124(input)
	
	var strSlice []string
	for _, val := range sequence {
		strSlice = append(strSlice, strconv.Itoa(val))
	}

	output := strings.Join(strSlice, "-")
	fmt.Println("Output:", output)
}
