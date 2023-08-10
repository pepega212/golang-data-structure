package main

import (
	"fmt"
	"strconv"
)

// func ArrayMerge(arrayA, arrayB []string) []string {
// 	merged := append(arrayA, arrayB...)
// 	return merged
// }

func Mapping(slice []string) map[string]int {
	countMap := make(map[string]int)

	for _, s := range slice {
		countMap[s]++
	}
	return countMap
}

func munculSekali(angka string) []int {
	countMap := make(map[int]int)

	for _, char := range angka {
		digit, _ := strconv.Atoi(string(char))
		countMap[digit]++
	}

	var result []int
	for digit, count := range countMap {
		if count == 1 {
			result = append(result, digit)
		}
	}

	return result
}

func main() {
	// fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// fmt.Println(ArrayMerge([]string{}, []string{}))

	// fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	// fmt.Println(Mapping([]string{}))

	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
