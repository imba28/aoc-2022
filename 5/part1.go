package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	errStackEmpty = errors.New("stack is empty")
)

type stack[T any] []T

func (s *stack[T]) pop() (T, error) {
	var emptyValue T
	if len(*s) == 0 {
		return emptyValue, errStackEmpty
	}
	idx := len(*s) - 1
	item := (*s)[idx]
	*s = (*s)[:idx]
	return item, nil
}

func (s *stack[T]) peek() (T, error) {
	var emptyValue T
	if len(*s) == 0 {
		return emptyValue, errStackEmpty
	}
	return (*s)[len(*s)-1], nil
}

func (s *stack[T]) push(r T) {
	*s = append(*s, r)
}

func rearrangeStack(stacks []stack[rune], repeat, from, to int) {
	for i := 0; i < repeat; i++ {
		v, err := stacks[from].pop()
		if errors.Is(errStackEmpty, err) {
			panic(fmt.Sprintf("Cannot pop empty stack %d", from))
		}

		stacks[to].push(v)
	}
}

func main() {
	r := regexp.MustCompile("^move (\\d+) from (\\d+) to (\\d+)$")
	stacks := []stack[rune]{
		{'B', 'Z', 'T'},
		{'V', 'H', 'T', 'D', 'N'},
		{'B', 'F', 'M', 'D'},
		{'T', 'J', 'G', 'W', 'V', 'Q', 'L'},
		{'W', 'D', 'G', 'P', 'V', 'F', 'Q', 'M'},
		{'V', 'Z', 'Q', 'G', 'H', 'F', 'S'},
		{'Z', 'S', 'N', 'R', 'L', 'T', 'C', 'W'},
		{'Z', 'H', 'W', 'D', 'J', 'N', 'R', 'M'},
		{'M', 'Q', 'L', 'F', 'D', 'S'},
	}

	f, _ := os.Open("input")
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		match := r.FindStringSubmatch(s.Text())
		repeat, from, to := mustParseInt(match[1]), mustParseInt(match[2])-1, mustParseInt(match[3])-1
		rearrangeStack(stacks, repeat, from, to)
	}

	for i := range stacks {
		v, err := stacks[i].peek()
		if errors.Is(errStackEmpty, err) {
			panic(err)
		}
		fmt.Print(string(v))
	}
}

func mustParseInt(v string) int {
	n, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return n
}
