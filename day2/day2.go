package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	infile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()
	scanner := bufio.NewScanner((infile))

	game_condition := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	game_result := 0
	game_result_2 := 0

	for scanner.Scan() {
		current_in := scanner.Text()
		current_result := meet_conditions(current_in, game_condition)
		current_result_2 := find_condition(current_in)
		game_result_2 += current_result_2
		game_result += current_result
	}
	fmt.Printf("part 1 :")
	fmt.Println(game_result)
	fmt.Printf("part 2 : ")
	fmt.Println(game_result_2)

}

func meet_conditions(xs string, ms map[string]int) int {
	game_num, _ := strconv.Atoi(strings.Split(strings.Split(xs, ":")[0], " ")[1])
	game_content := strings.Split(strings.Split(xs, ":")[1], ";")
	meet := true
	for _, s := range game_content {
		current_vs := strings.Split(s, ",")
		for _, s_ := range current_vs {
			current_hs := strings.Split(strings.TrimLeft(s_, " "), " ")
			for i, h := range current_hs {
				if _, err := strconv.Atoi(h); err != nil {
					v, _ := strconv.Atoi(current_hs[i-1])
					if ms[h] >= v {
						meet = meet && true
					} else {
						meet = meet && false
					}
				}

			}
		}
	}
	if meet {
		return game_num
	} else {
		return 0
	}

}

func find_condition(xs string) int {
	game_results := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	game_content := strings.Split(strings.Split(xs, ":")[1], ";")
	for _, s := range game_content {
		current_vs := strings.Split(s, ",")
		for _, s_ := range current_vs {
			current_hs := strings.Split(strings.TrimLeft(s_, " "), " ")
			for i, h := range current_hs {
				if _, err := strconv.Atoi(h); err != nil {
					v, _ := strconv.Atoi(current_hs[i-1])
					if game_results[h] < v {
						game_results[h] = v
					}
				}

			}
		}
	}
	result := 1
	for _, v := range game_results {
		result *= v
	}
	return result

}
