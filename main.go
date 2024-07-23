package main

import (
	"fmt"
	"math/rand"
)

func main() {
	play()
}

type Nim struct {
	Board  []int
	Player int
}

func NewNim() *Nim {
	return &Nim{Board: []int{1, 3, 5, 7}, Player: 0}
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
	return pile > 0 && count > 0 && pile <= len(n.Board) && count <= n.Board[pile-1] && n.Board[pile-1] > 0
}

func (n *Nim) SwitchPlayer() {
	n.Player = n.Player ^ 1
}

func (n *Nim) AvailableMoves() [][2]int {
	var resp [][2]int
	for pile, countObjects := range n.Board {
		for j := 1; j <= countObjects; j++ {
			resp = append(resp, [2]int{pile + 1, j})
		}
	}
	return resp
}

func play() {
	humanPlayer := rand.Intn(2)

	nim := NewNim()

	for {

		// print the board
		fmt.Println()
		for i, countObjects := range nim.Board {
			fmt.Printf("Pile %d: %d\n", i+1, countObjects)
		}

		var pile, count int

		if nim.Player == humanPlayer {
			for {
				fmt.Println()
				fmt.Println("Human's turn")
				fmt.Print("Choose Pile: ")
				fmt.Scan(&pile)
				fmt.Print("Choose Count: ")
				fmt.Scan(&count)
				if nim.IsValidMove(pile, count) {
					break
				}
				fmt.Println("Invalid move, try again")
			}
		} else {
			fmt.Println()
			fmt.Println("AI's turn")
			moves := nim.AvailableMoves()
			move := moves[rand.Intn(len(moves))]
			pile = move[0]
			count = move[1]
			fmt.Printf("AI chose to take %d from pile %d.\n", count, pile)
		}

		// make a move
		nim.Move(pile, count)

		// shift player
		nim.SwitchPlayer()

		// check game over
		gameOver := nim.GameOver()

		// finish the game
		if gameOver {
			winner := "AI"
			if nim.Player == humanPlayer {
				winner = "human"
			}
			fmt.Println()
			fmt.Println("GAME OVER")
			fmt.Printf("Winners is %s\n", winner)
			break
		}
	}

}
