package tictactoe

import (
	"testing"
)

func Test_Minimax(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PX, PX, PE,
		PE, PE, PE,
		PO, PE, PO}
	b.isXTurn = true
	best := b.CalcBestMove()
	if best.target != 2 {
		t.Error(best)
	}
}
