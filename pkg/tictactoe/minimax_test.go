package tictactoe

import (
	"testing"
)

func Test_Minimax(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PX, PO, PE,
		PX, PX, PE,
		PO, PE, PO}
	b.isXTurn = true
	best := b.CalcBestMove()
	if best.target != 5 {
		t.Error(best)
	}
}
