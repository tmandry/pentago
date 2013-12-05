package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	pg "./pentago"
)

var colors = map[pg.Piece]string{pg.White: "White", pg.Black: "Black"}

const (
	human = iota
	random
	ai
)

func main() {
	strategies := map[string]int{"human": human, "random": random, "ai": ai}
	flag.Parse()

	whiteStrategy, blackStrategy := ai, ai
	args, ok := flag.Args(), true
	if len(args) > 0 {
		whiteStrategy, ok = strategies[args[0]]
	}
	if len(args) > 1 && ok {
		blackStrategy, ok = strategies[args[1]]
	}
	if !ok {
		fmt.Println("Usage: pentago [white-strategy] [black-strategy]")
		fmt.Println("Where each strategy is one of: human, random, ai")
		os.Exit(1)
	}

	game := pg.NewGame()
	play(&game, whiteStrategy, blackStrategy)
}

func play(game *pg.Game, whiteStrategy, blackStrategy int) {
	strategy := map[pg.Piece]int{pg.White: whiteStrategy, pg.Black: blackStrategy}

	for {
		// Print and check game state
		fmt.Print(game.Board)
		if winner := game.CheckWinner(); winner != pg.Empty {
			fmt.Printf("%s won!", colors[winner])
			break
		}
		fmt.Printf("\n%s's move\n", colors[game.Turn])

		// Execute move according to strategy for this turn
		var move pg.Move
		start := time.Now()
		switch strategy[game.Turn] {
		case human:
			move = promptForMove(game)
		case random:
			move = game.RandomMove()
		case ai:
			move = game.BestMove(game.Turn)
		}
		fmt.Printf("finished in %v\n", time.Since(start))

		// Check for errors
		if ok := game.Move(move); !ok {
			fmt.Println("Error")
		}
		fmt.Println(move)
	}
}

func promptForMove(game *pg.Game) (move pg.Move) {
	var r, c, sub, dir int
	for ok := false; !ok; {
		fmt.Printf("%s> ", colors[game.Turn])

		_, err := fmt.Scanf("%d %d %d %d", &r, &c, &sub, &dir)
		switch {
		case err != nil:
			fmt.Println("Invalid format")
		case r < 0 || r >= 6 || c < 0 || c >= 6:
			fmt.Println("Row and columns are 0-5")
		case sub < 0 || sub >= 4:
			fmt.Println("Subboard is 0-3")
		case !(dir == 0 || dir == 1):
			fmt.Println("Direction is 0 for clockwise, 1 for counterclockwise")
		default:
			move = pg.NewMove(r, c, sub, dir)
			if ok = move.IsValid(game.Board); !ok {
				fmt.Println("Invalid move")
			}
		}
	}

	return move
}
