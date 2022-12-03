package main

import "os"
import "strings"
import "bufio"
import "container/ring"
import "fmt"

const (
	rock = iota
	scissor
	paper
)

const (
	win = iota
	loss
	draw
)

type gameMove int

func (g gameMove) score(roundResult int) int {
	var score int
	switch g {
	case rock:
		score = 1
	case scissor:
		score = 3
	case paper:
		score = 2
	}

	if roundResult == win {
		score += 6
	} else if roundResult == draw {
		score += 3
	}

	return score
}

func newGameMove(r rune) gameMove {
	switch r {
	case 'X':
		fallthrough
	case 'A':
		return gameMove(rock)
	case 'Y':
		fallthrough
	case 'B':
		return gameMove(paper)
	case 'Z':
		fallthrough
	case 'C':
		return gameMove(scissor)
	default:
		panic("invalid game move")
	}
}

type game struct {
	list *ring.Ring
}

func (g game) beats(a, b gameMove) bool {
	in := g.list
	if in.Value.(int) == int(a) {
		return in.Next().Value.(int) == int(b)
	}

	for p := in.Next(); p != in; p = p.Next() {
		if p.Value.(int) == int(a) {
			return p.Next().Value.(int) == int(b)
		}
	}

	return false
}

func (g game) draw(a, b gameMove) bool {
	return a == b
}

func (g game) scoreOf(a, b gameMove) int {
	if g.draw(a, b) {
		return b.score(draw)
	}
	if g.beats(a, b) {
		return b.score(loss)
	}
	return b.score(win)
}

func (g game) play() int {
	f, _ := os.Open("input")
	defer f.Close()

	s := bufio.NewScanner(f)
	var score int
	for s.Scan() {
		chars := strings.Split(s.Text(), "")
		moveA, moveB := newGameMove(rune(chars[0][0])), newGameMove(rune(chars[2][0]))

		score += g.scoreOf(moveA, moveB)
	}

	return score
}

func newGame() game {
	gameItems := []int{
		rock,
		scissor,
		paper,
	}
	g := game{
		list: ring.New(len(gameItems)),
	}

	for i := range gameItems {
		g.list.Value = gameItems[i]
		g.list = g.list.Next()
	}

	return g
}

func main() {
	g := newGame()

	fmt.Println("Score:", g.play())
}
