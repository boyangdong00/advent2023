package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	scanner := bufio.NewScanner((infile))

	w_nums := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	w_numsr := []string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}

	result := 0

	for scanner.Scan() {
		current_in := scanner.Text()
		left_num := findFirstInt(current_in, w_nums)
		right_num := findFirstInt(reverse(current_in), w_numsr)
		t1, _ := contains(w_nums, left_num)
		t2, _ := contains(w_numsr, right_num)
		if t1 != -1 && len(left_num) != 1 {
			left_num = strconv.Itoa(t1 + 1)
		}
		if t2 != -1 && len(right_num) != 1 {
			right_num = strconv.Itoa(t2 + 1)

		}

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

func findFirstInt(s string, ws []string) string {
	w_num := ""
	for _, char := range s {
		if unicode.IsDigit(char) {
			return string(char)
		}
		w_num = w_num + string(char)
		t, s := contains(ws, w_num)
		if t != -1 {
			return s
		}
	}
	return ""
}

func contains(ws []string, s string) (int, string) {
	for i, c := range ws {
		if strings.Contains(s, c) {
			return i, c
		}
	}
	return -1, ""
}
