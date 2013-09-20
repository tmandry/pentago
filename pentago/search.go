package pentago

import "math/rand"

func (b Board) RandomMove() Move {
	moves := b.ValidMoves()
	return moves[rand.Intn(len(moves))]
}
