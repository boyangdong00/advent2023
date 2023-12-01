package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	scanner := bufio.NewScanner((infile))

	result := 0

	for scanner.Scan() {
		current_in := scanner.Text()
		left_num := findFirstInt(current_in)
		right_num := findFirstInt(reverse(current_in))
		current := 0
		if left_num != "" {
			current, _ = strconv.Atoi(left_num + right_num)
		}
		result = result + current
	}
	fmt.Println(result)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func findFirstInt(s string) string {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return string(char)
		}
	}
	return ""
}
