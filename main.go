package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

type Nim struct {
	Board []int
}

func NewNim() *Nim {
	return &Nim{Board: []int{1, 3, 5, 7}}
}

func (n *Nim) Move(pile, count int) {
	n.Board[pile-1] = n.Board[pile-1] - count
}

func (n *Nim) GameOver() bool {
	for _, countObjects := range n.Board {
		if countObjects != 0 {
			return false
		}
	}
	return true
}

func (n *Nim) IsValidMove(pile, count int) bool {
	return pile > 0 && pile <= len(n.Board) && count <= n.Board[pile-1] && n.Board[pile-1] > 0
}

func play() {
	humanPlayer := rand.Intn(2)

	nim := NewNim()

	for {

		// print the board
		fmt.Println()
		fmt.Printf("Player %d turn\n", humanPlayer)
		for i, countObjects := range nim.Board {
			fmt.Printf("Pile %d: %d\n", i+1, countObjects)
		}

		var pile, count int

		// ask a move
		for {
			fmt.Print("Choose Pile: ")
			fmt.Scan(&pile)
			fmt.Print("Choose Count: ")
			fmt.Scan(&count)
			if nim.IsValidMove(pile, count) {
				break
			}
			fmt.Println("Invalid move, try again")
			fmt.Println()
		}

		// make a move
		nim.Move(pile, count)

		// shift player
		humanPlayer = humanPlayer ^ 1

		// check winner
		winner := nim.GameOver()

		// finish the game
		if winner {
			fmt.Println()
			fmt.Printf("Winners is Player %d\n", humanPlayer)
			break
		}
	}

}
