package main

import "fmt"

type WinStrategy interface {
	CheckWinner(board [][]string) string
}

type RowWinStrategy struct{}

func (r *RowWinStrategy) CheckWinner(board [][]string) string {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
	}
	return ""
}

type ColumnWinStrategy struct{}

func (c *ColumnWinStrategy) CheckWinner(board [][]string) string {
	for i := 0; i < 3; i++ {
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	return ""
}

type DiagonalWinStrategy struct{}

func (d *DiagonalWinStrategy) CheckWinner(board [][]string) string {
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return ""
}

type WinnerChecker struct {
	strategies []WinStrategy
}

// AddStrategy adds a new strategy to the checker
func (w *WinnerChecker) AddStrategy(strategy WinStrategy) {
	w.strategies = append(w.strategies, strategy)
}

// GetWinner checks all strategies for a winner
func (w *WinnerChecker) GetWinner(board [][]string) string {
	for _, strategy := range w.strategies {
		if winner := strategy.CheckWinner(board); winner != "" {
			return winner
		}
	}
	return ""
}

func main() {
	// Initialize the board
	board := [][]string{
		{"X", "X", "O"},
		{"O", "", "O"},
		{"", "", "X"},
	}

	// Initialize WinnerChecker with strategies
	checker := &WinnerChecker{}
	checker.AddStrategy(&RowWinStrategy{})
	checker.AddStrategy(&ColumnWinStrategy{})
	checker.AddStrategy(&DiagonalWinStrategy{})

	// Check for a winner
	winner := checker.GetWinner(board)
	if winner != "" {
		fmt.Printf("Winner: %s\n", winner)
	} else {
		fmt.Println("No winner yet.")
	}
}
