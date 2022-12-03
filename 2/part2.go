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
	case 'A':
		return gameMove(rock)
	case 'B':
		return gameMove(paper)
	case 'C':
		return gameMove(scissor)
	default:
		panic("invalid game move")
	}
}

func expectedResult(r rune) int {
	switch r {
	case 'X':
		return loss
	case 'Y':
		return draw
	case 'Z':
		return win
	default:
		panic("invalid round result")
	}
}

type game struct {
	list *ring.Ring
}

func (g game) findGameItem(a gameMove) *ring.Ring {
	in := g.list
	if in.Value.(gameMove) == a {
		return in
	}

	for p := in.Next(); p != in; p = p.Next() {
		if p.Value.(gameMove) == a {
			return p
		}
	}

	return nil
}

func (g game) draw(a, b gameMove) bool {
	return a == b
}

func (g game) scoreOf(a gameMove, result int) int {
	var b *ring.Ring
	if result == draw {
		return a.score(draw)
	}

	item := g.findGameItem(a)
	if result == win {
		b = item.Prev()
	} else {
		b = item.Next()
	}

	return b.Value.(gameMove).score(result)
}

func (g game) play() int {
	f, _ := os.Open("input")
	defer f.Close()

	s := bufio.NewScanner(f)
	var score int
	for s.Scan() {
		chars := strings.Split(s.Text(), "")
		move, result := newGameMove(rune(chars[0][0])), expectedResult(rune(chars[2][0]))

		score += g.scoreOf(move, result)
	}

	return score
}

func newGame() game {
	gameItems := []gameMove{
		gameMove(rock),
		gameMove(scissor),
		gameMove(paper),
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
