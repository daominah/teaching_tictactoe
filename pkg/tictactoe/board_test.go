package tictactoe

import (
	"testing"
)

func TestBoard_String(t *testing.T) {
	b := NewBoard()
	b.squares[2] = PO
	b.squares[4] = PX
}

func TestBoard_CalcLegalMoves(t *testing.T) {
	b := NewBoard()
	moves := b.CalcLegalMoves()
	if len(moves) != 9 {
		t.Error()
	}
	b.squares[4] = PX
	moves = b.CalcLegalMoves()
	if len(moves) != 8 {
		t.Error()
	}
}

func TestBoard_CheckResult(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PX, PX, PX,
		PO, PO, PE,
		PE, PE, PE}
	if b.CheckResult() != Win {
		t.Error()
	}
	b.squares = []Piece{
		PO, PX, PX,
		PX, PO, PE,
		PX, PE, PO}
	if b.CheckResult() != Loss {
		t.Error()
	}
	b.squares = []Piece{
		PO, PX, PX,
		PX, PO, PO,
		PX, PO, PX}
	if b.CheckResult() != Draw {
		t.Error()
	}
	b.squares = []Piece{
		PO, PO, PX,
		PX, PE, PO,
		PX, PO, PX}
	if b.CheckResult() != Playing {
		t.Error()
	}
}