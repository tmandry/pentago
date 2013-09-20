package pentago

import (
	"bytes"
)

type Piece int

const (
	Empty Piece = iota
	Black
	White
)

type Board [][]Piece

func NewBoard() Board {
	b := make([][]Piece, 6)
	for r := 0; r < 6; r++ {
		b[r] = make([]Piece, 6)
	}
	return b
}

func (b Board) String() string {
	var buffer bytes.Buffer
	for _, row := range b {
		for _, piece := range row {
			switch piece {
			case Empty:
				buffer.WriteString(" .")
			case Black:
				buffer.WriteString(" B")
			case White:
				buffer.WriteString(" W")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

type Coord struct {
	Row int
	Col int
}

func (b Board) GetPiece(c Coord) Piece {
	return b[c.Row][c.Col]
}

func (b Board) getTriple(start Coord, deltaR int, deltaC int) [3]Piece {
	return [3]Piece{
		b[start.Row][start.Col],
		b[start.Row+deltaR][start.Col+deltaC],
		b[start.Row+2*deltaR][start.Col+2*deltaC],
	}
}

func (b Board) setTriple(start Coord, deltaR int, deltaC int, triple [3]Piece) {
	b[start.Row][start.Col] = triple[0]
	b[start.Row+deltaR][start.Col+deltaC] = triple[1]
	b[start.Row+2*deltaR][start.Col+2*deltaC] = triple[2]
}

const (
	Clockwise = iota
	CounterClockwise
)

func (b Board) Rotate(subRow int, subCol int, direction int) {
	topLeft := Coord{3*subRow, 3*subCol}
	botLeft := Coord{topLeft.Row+2, topLeft.Col}
	botRight := Coord{botLeft.Row, botLeft.Col+2}
	topRight := Coord{topLeft.Row, topLeft.Col+2}

	top := b.getTriple(topLeft, 0, 1)
	switch direction {
	case Clockwise:
		b.setTriple(topLeft, 0, 1, b.getTriple(botLeft, -1, 0))
		b.setTriple(botLeft, -1, 0, b.getTriple(botRight, 0, -1))
		b.setTriple(botRight, 0, -1, b.getTriple(topRight, 1, 0))
		b.setTriple(topRight, 1, 0, top)
	case CounterClockwise:
		b.setTriple(topLeft, 0, 1, b.getTriple(topRight, 1, 0))
		b.setTriple(topRight, 1, 0, b.getTriple(botRight, 0, -1))
		b.setTriple(botRight, 0, -1, b.getTriple(botLeft, -1, 0))
		b.setTriple(botLeft, -1, 0, top)
	}
}

// Checks if a span of consecutive pieces of the same color exists on a given
// row or column. If it does, returns the color of the span. If not, returns
// Empty.
func (b Board) checkSpan(coord Coord, deltaR int, deltaC int, length int) Piece {
	// If first two are the same, they are part of the span in this row if it
	// exists. If not, the only way is if the span starts at index 1.
	if b[coord.Row][coord.Col] != b[coord.Row+deltaR][coord.Col+deltaC] {
		coord.Row += deltaR
		coord.Col += deltaC
	}

	color := b[coord.Row][coord.Col]
	for i := 1; i < length; i++ {
		coord.Row += deltaR
		coord.Col += deltaC
		if b[coord.Row][coord.Col] != color {
			return Empty
		}
	}
	return color
}

// Returns the color of the winner if there is one, or Empty if none.
func (b Board) CheckWinner() Piece {
	// Check rows
	for r := 0; r < 6; r++ {
		if color := b.checkSpan(Coord{r, 0}, 0, 1, 5); color != Empty {
			return color
		}
	}
	// Check cols
	for c := 0; c < 6; c++ {
		if color := b.checkSpan(Coord{0, c}, 1, 0, 5); color != Empty {
			return color
		}
	}
	return Empty
}

type Game struct {
	Board
	Turn Piece
}

func NewGame() Game {
	return Game{
		Board: NewBoard(),
		Turn: White,
	}
}

func (g *Game) Move(r, c, sub, dir int) bool {
	if g.Board[r][c] != Empty {
		return false
	}

	g.Board[r][c] = g.Turn
	g.Board.Rotate(sub / 2, sub % 2, dir)
	switch g.Turn {
	case White: g.Turn = Black
	case Black: g.Turn = White
	}
	return true
}
