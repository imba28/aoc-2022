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

func sumSlice(s []int) int {
	var vv int
	for _, v := range s {
		vv += v
	}
	return vv
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

	fmt.Println("Part 1", values[0])

	fmt.Println("Part 2", sumSlice(values[:3]))
}
