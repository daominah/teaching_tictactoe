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

func Test_Minimax2(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PE, PO, PE,
		PE, PX, PE,
		PE, PE, PE}
	b.isXTurn = true
	best := b.CalcBestMove()
	if best.target != 0 && best.target != 2 {
		t.Error(best)
	}
}

func Test_Minimax3(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PE, PE, PE,
		PE, PX, PE,
		PE, PE, PE}
	b.isXTurn = true
	best := b.CalcBestMove()
	if best.target != 0 && best.target != 2 &&
		best.target != 6 && best.target != 8 {
		t.Error(best)
	}
}
