package main

import (
	"fmt"

	pg "./pentago"
)

func main() {
	game := pg.NewGame()
	playAIvAI(&game)
}

var colors = map[pg.Piece]string{pg.White: "White", pg.Black: "Black"}

func playHvH(game *pg.Game) {
	for {
		fmt.Print(game.Board)
		if winner := game.CheckWinner(); winner != pg.Empty {
			fmt.Printf("%s won!", colors[winner])
			break
		}
		move := promptForMove(game)
		game.Move(move)
		fmt.Println(move)
	}
}

func playHvR(game *pg.Game) {
	for {
		fmt.Print(game.Board)
		if winner := game.CheckWinner(); winner != pg.Empty {
			fmt.Printf("%s won!", colors[winner])
			break
		}

		var move pg.Move
		if game.Turn == pg.White {
			move = promptForMove(game)
		} else {
			move = game.RandomMove()
		}

		if ok := game.Move(move); !ok {
			fmt.Println("Error")
		}
		fmt.Println(move)
	}
}

func playHvAI(game *pg.Game) {
	for {
		fmt.Print(game.Board)
		if winner := game.CheckWinner(); winner != pg.Empty {
			fmt.Printf("%s won!", colors[winner])
			break
		}

		var move pg.Move
		if game.Turn == pg.White {
			move = promptForMove(game)
		} else {
			move = game.BestMove(pg.Black)
		}

		if ok := game.Move(move); !ok {
			fmt.Println("Error")
		}
		fmt.Println(move)
	}
}

func playRvAI(game *pg.Game) {
	for {
		fmt.Print(game.Board)
		if winner := game.CheckWinner(); winner != pg.Empty {
			fmt.Printf("%s won!", colors[winner])
			break
		}

		var move pg.Move
		if game.Turn == pg.White {
			move = game.RandomMove()
		} else {
			move = game.BestMove(pg.Black)
		}

		if ok := game.Move(move); !ok {
			fmt.Println("Error")
		}
		fmt.Println(move)
	}
}

func playAIvAI(game *pg.Game) {
	for {
		fmt.Print(game.Board)
		if winner := game.CheckWinner(); winner != pg.Empty {
			fmt.Printf("%s won!", colors[winner])
			break
		}

		var move pg.Move
		if game.Turn == pg.White {
			move = game.BestMove(pg.White)
		} else {
			move = game.BestMove(pg.Black)
		}

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
