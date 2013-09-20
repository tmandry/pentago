package pentago

import (
	"fmt"
	"testing"
)

func startBoard() Board {
	b := NewBoard()
	b[0][1] = White
	b[0][2] = Black
	b[0][4] = White
	b[1][1] = Black
	b[1][4] = Black
	b[2][2] = White
	b[2][3] = Black
	b[3][0] = Black
	b[3][1] = White
	b[3][2] = Black
	b[3][5] = Black
	b[4][0] = White
	b[4][2] = Black
	b[4][5] = White
	b[5][0] = White
	b[5][2] = Black
	b[5][4] = White
	b[5][5] = White
	return b
}

func ExampleStartBoard() {
	fmt.Print(startBoard())
	// Output:
	//  . W B . W .
	//  . B . . B .
	//  . . W B . .
	//  B W B . . B
	//  W . B . . W
	//  W . B . W W
}

func ExampleRotateCCW() {
	b := startBoard()
	b.Rotate(1, 0, CounterClockwise)
	fmt.Print(b)
	// Output:
	//  . W B . W .
	//  . B . . B .
	//  . . W B . .
	//  B B B . . B
	//  W . . . . W
	//  B W W . W W
}

func ExampleRotateCW() {
	b := startBoard()
	b.Rotate(0, 1, Clockwise)
	fmt.Print(b)
	// Output:
	//  . W B B . .
	//  . B . . B W
	//  . . W . . .
	//  B W B . . B
	//  W . B . . W
	//  W . B . W W
}

func TestCheckWinner(t *testing.T) {
	b := NewBoard()
	if v := b.CheckWinner(); v != Empty {
		t.Errorf("want Empty, got %v", v)
	}

	// Test rows
	b[0] = []Piece{White, White, White, White, White, Empty}
	if v := b.CheckWinner(); v != White {
		t.Errorf("want White, got %v", v)
	}
	b[0] = []Piece{Black, Black, Black, Black, Black, Empty}
	if v := b.CheckWinner(); v != Black {
		t.Errorf("want Black, got %v", v)
	}

	// Test columns
	b = NewBoard()
	b[2][1], b[2][2], b[2][3], b[2][4], b[2][5] = White, White, White, White, White
	if v := b.CheckWinner(); v != White {
		t.Errorf("want White, got %v", v)
	}
	b[2][1], b[2][2], b[2][3], b[2][4], b[2][5] = White, White, Black, White, White
	if v := b.CheckWinner(); v != Empty {
		t.Errorf("want Empty, got %v", v)
	}
}

func ExampleMoveRotatingDifferentSubboard() {
	g := NewGame()
	g.Board = startBoard()
	g.Turn = White
	ok := g.Move(2, 1, 2, CounterClockwise)
	fmt.Println(ok)
	fmt.Print(g)
	// Output:
	// true
	//  . W B . W .
	//  . B . . B .
	//  . W W B . .
	//  B B B . . B
	//  W . . . . W
	//  B W W . W W
}

func ExampleMoveRotatingSameSubboard() {
	g := NewGame()
	g.Board = startBoard()
	g.Turn = White
	ok := g.Move(5, 1, 2, CounterClockwise)
	fmt.Println(ok)
	fmt.Print(g)
	// Output:
	// true
	//  . W B . W .
	//  . B . . B .
	//  . . W B . .
	//  B B B . . B
	//  W . W . . W
	//  B W W . W W
}

func ExampleInvalidMove() {
	g := NewGame()
	g.Board = startBoard()
	g.Turn = White
	ok := g.Move(0, 1, 2, CounterClockwise)
	fmt.Println(ok)
	fmt.Print(g)
	// Output:
	// false
	//  . W B . W .
	//  . B . . B .
	//  . . W B . .
	//  B W B . . B
	//  W . B . . W
	//  W . B . W W
}

func TestMoveAlternatesTurns(t *testing.T) {
	g := NewGame()
	g.Board = startBoard()
	g.Turn = White
	_ = g.Move(5, 1, 2, CounterClockwise)
	if g.Turn != Black {
		t.Error("Expected Turn to switch to black")
	}
	_ = g.Move(5, 3, 2, CounterClockwise)
	if g.Turn != White {
		t.Error("Expected Turn to switch to white")
	}
	_ = g.Move(5, 3, 2, CounterClockwise)
	if g.Turn != White {
		t.Error("Expected Turn to stay the same for invalid move")
	}
}
