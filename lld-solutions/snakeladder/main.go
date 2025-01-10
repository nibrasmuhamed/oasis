package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Game represents the snake and ladder game
type Game struct {
	BoardSize int
	Board     []int // Represents each square on the board, where 0 means no snake or ladder, positive means ladder to that position, negative means snake to that position
	Players   []*Player
}

// Player represents a player in the game
type Player struct {
	Name     string
	Position int
}

// NewGame creates a new snake and ladder game with a board of specified size
func NewGame(size int, numPlayers int) *Game {
	game := &Game{
		BoardSize: size,
		Board:     make([]int, size+1), // +1 because board indices start from 1
		Players:   make([]*Player, numPlayers),
	}

	// Initialize players
	for i := 0; i < numPlayers; i++ {
		game.Players[i] = &Player{
			Name:     fmt.Sprintf("Player %d", i+1),
			Position: 1, // Start from square 1
		}
	}

	return game
}

// SetSnakeOrLadder sets a snake or ladder on the board
func (g *Game) SetSnakeOrLadder(start, end int) {
	if start > 0 && start <= g.BoardSize && end > 0 && end <= g.BoardSize {
		if end > start {
			g.Board[start] = end - start // Positive number represents ladder
		} else {
			g.Board[start] = end - start // Negative number represents snake
		}
	}
}

// RollDice simulates rolling a 6-sided dice
func RollDice() int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return random.Intn(6) + 1
}

// PlayRound simulates one round of the game
func (g *Game) PlayRound(player *Player) {
	diceRoll := RollDice()
	newPosition := player.Position + diceRoll

	if newPosition <= g.BoardSize {
		if g.Board[newPosition] != 0 {
			newPosition += g.Board[newPosition]
			if newPosition < 1 {
				newPosition = 1
			}
		}
		player.Position = newPosition
	}
}

// IsGameFinished checks if any player has reached the last square
func (g *Game) IsGameFinished() bool {
	for _, player := range g.Players {
		if player.Position == g.BoardSize {
			return true
		}
	}
	return false
}

func main() {
	game := NewGame(100, 3)       // Create a game board of size 100 with 2 players
	game.SetSnakeOrLadder(14, 4)  // Set a snake from 14 to 4
	game.SetSnakeOrLadder(26, 8)  // Set a snake from 26 to 8
	game.SetSnakeOrLadder(38, 14) // Set a snake from 38 to 14
	game.SetSnakeOrLadder(50, 34) // Set a snake from 50 to 34
	game.SetSnakeOrLadder(76, 42) // Set a snake from 76 to 42
	game.SetSnakeOrLadder(91, 73) // Set a snake from 91 to 73
	game.SetSnakeOrLadder(3, 17)  // Set a ladder from 3 to 17
	game.SetSnakeOrLadder(10, 25) // Set a ladder from 10 to 25
	game.SetSnakeOrLadder(22, 51) // Set a ladder from 22 to 51
	game.SetSnakeOrLadder(43, 85) // Set a ladder from 43 to 85
	game.SetSnakeOrLadder(63, 95) // Set a ladder from 63 to 95
	game.SetSnakeOrLadder(71, 91) // Set a ladder from 71 to 91

	fmt.Println("Snake and Ladder Game Starts!")

	for !game.IsGameFinished() {
		for _, player := range game.Players {
			game.PlayRound(player)
			fmt.Printf("%s rolled dice and moved to square %d\n", player.Name, player.Position)
			if game.IsGameFinished() {
				fmt.Printf("%s wins!\n", player.Name)
				break
			}
		}
	}
}
