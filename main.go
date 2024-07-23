package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ai := NewNimAI(0.5, 0.1)
	ai.Train(100000)
	play(ai)
}

type Board []int

func (b Board) AvailableMoves() [][2]int {
	var resp [][2]int
	for pile, countObjects := range b {
		for j := 1; j <= countObjects; j++ {
			resp = append(resp, [2]int{pile + 1, j})
		}
	}
	return resp
}

func (b Board) Copy() Board {
	dst := make(Board, len(b))
	copy(dst, b)
	return dst
}

type Nim struct {
	Board  Board
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

func play(ai *NimAI) {
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
			move := ai.ChooseMove(nim.Board, false)
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

type NimAI struct {
	q       map[string]float64
	alpha   float64
	epsilon float64
	r       *rand.Rand
}

func NewNimAI(alpha, epsilon float64) *NimAI {
	return &NimAI{
		q:       make(map[string]float64),
		alpha:   alpha,
		epsilon: epsilon,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ai *NimAI) getQValue(state Board, move [2]int) float64 {
	hash := ai.hashAction(state, move)
	return ai.q[hash]
}

func (ai *NimAI) updateQValue(state Board, move [2]int, oldQ, reward, futureReward float64) {
	newQ := oldQ + ai.alpha*((reward+futureReward)-oldQ)
	hash := ai.hashAction(state, move)
	ai.q[hash] = newQ
}

func (ai *NimAI) hashAction(state Board, move [2]int) string {
	return fmt.Sprintf("%v|%v", state, move)
}

func (ai *NimAI) bestFutureReward(state Board) float64 {
	actions := state.AvailableMoves()
	if len(actions) == 0 {
		return 0
	}

	var bestActionValue float64 = 0

	for _, action := range actions {
		actionValue := ai.getQValue(state, action)
		if bestActionValue == 0 || actionValue > bestActionValue {
			bestActionValue = actionValue
		}
	}

	return bestActionValue
}

func (ai *NimAI) Update(oldState Board, move [2]int, newState Board, reward float64) {
	oldQ := ai.getQValue(oldState, move)
	bestFuture := ai.bestFutureReward(newState)
	ai.updateQValue(oldState, move, oldQ, reward, bestFuture)
}

func (ai *NimAI) ChooseMove(state Board, epsilon bool) [2]int {
	actions := state.AvailableMoves()

	if epsilon && ai.r.Float64() <= ai.epsilon {
		return actions[ai.r.Intn(len(actions))]
	}

	var (
		bestAction      [2]int
		bestActionValue float64
	)

	for _, action := range actions {
		actionValue := ai.getQValue(state, action)
		if bestActionValue == 0 || actionValue > bestActionValue {
			bestActionValue = actionValue
			bestAction = action
		}
	}

	return bestAction

}

func (ai *NimAI) Train(n int) {

	type action struct {
		state Board
		move  [2]int
	}

	for range n {
		game := NewNim()
		lastMoves := map[int]action{}

		for {
			currentState := game.Board.Copy()
			move := ai.ChooseMove(currentState, true)
			lastMoves[game.Player] = action{state: currentState, move: move}

			game.Move(move[0], move[1])
			game.SwitchPlayer()
			newState := game.Board.Copy()

			if game.GameOver() {
				// current move cause the lost
				ai.Update(currentState, move, newState, -1)
				// last move cause the win
				ai.Update(lastMoves[game.Player].state, lastMoves[game.Player].move, newState, 1)
				break
			}

			if len(lastMoves[game.Player].state) > 0 {
				ai.Update(lastMoves[game.Player].state, lastMoves[game.Player].move, newState, 0)
			}
		}
	}
}
