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
	if best.Target != 5 {
		t.Error(best)
	}
}

func Test_Minimax2(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PX, PE, PE,
		PE, PE, PE,
		PE, PE, PE}
	b.isXTurn = false
	best := b.CalcBestMove()
	if best.Target != 4 {
		t.Error(best)
	}
}

func Test_Minimax3(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PE, PE, PE,
		PE, PX, PE,
		PE, PE, PE}
	b.isXTurn = false
	best := b.CalcBestMove()
	if best.Target != 0 && best.Target != 2 &&
		best.Target != 6 && best.Target != 8 {
		t.Error(best)
	}
}
