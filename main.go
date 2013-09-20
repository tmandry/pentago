package main

import (
	"fmt"

	pg "./pentago"
)

func main() {
	game := pg.NewGame()
	play(game)
}

func play(game pg.Game) {
	directions := map[int]string{pg.Clockwise: "CW", pg.CounterClockwise: "CCW"}
	colors := map[pg.Piece]string{pg.White: "White", pg.Black: "Black"}

	for {
		fmt.Print(game.Board)
		if winner := game.CheckWinner(); winner != pg.Empty {
			fmt.Printf("%s won!", colors[winner])
			break
		}

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
				ok = game.Move(r, c, sub, dir)
				if !ok {
					fmt.Println("Invalid move")
				}
			}
		}

		fmt.Printf("Put piece (%d, %d), rotate subboard %d %s\n", r, c, sub, directions[dir])
	}
}
