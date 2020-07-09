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

func TestBoard_MakeMove_TakeBack(t *testing.T) {
	b := NewBoard()
	b.MakeMove(Move{target: 4})
	b.MakeMove(Move{target: 1})
	if len(b.history) != 2 {
		t.Errorf("wrong history: %v", b.history)
	}
	b.TakeBack()
	if len(b.history) != 1 {
		t.Errorf("wrong history: %v", b.history)
	}
	if b.squares[1] != PE || b.squares[4] != PX {
		t.Errorf("wrong TakeBack")
	}
	b.TakeBack()
	if b.squares[4] != PE {
		t.Errorf("wrong TakeBack")
	}
}

func TestBoard_MakeRandomMove(t *testing.T) {
	for nMatch := 0; nMatch < 10; nMatch++ {
		b := NewBoard()
		for {
			bestMove, err := b.CalcRandomMove()
			if err != nil {
				t.Fatal(err, b)
			}
			isValid := b.MakeMove(bestMove)
			if !isValid {
				t.Fatalf("board: %v, bestMove: %v", b, bestMove)
			}
			result := b.CheckResult()
			//t.Log(b)
			if result != Playing {
				//t.Logf("___________________________________________ %v", result)
				break
			}
		}
	}
}
