package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func main() {
	f, _ := os.Open("input")
	defer f.Close()

	s := bufio.NewScanner(f)
	var values []int
	var counter int
	for s.Scan() {
		line := s.Text()
		if line == "" {
			values = append(values, counter)
			counter = 0
			continue
		}
		counter += toInt(line)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	fmt.Println(values[0])
}
