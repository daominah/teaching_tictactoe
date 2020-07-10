package tictactoe

import (
	"testing"

	"github.com/daominah/teaching_tictactoe/pkg/minimax"
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

	for i := 0; i < 5; i++ {
		stats := minimax.NewStats()
		minimax.Minimax(b, stats, 999)
		if i == 0 {
			t.Logf("nNodes: %v", stats.NNodes)
		}
		bm, _ := stats.PosTable[b.Hash()].BestMove.(Move)
		//t.Logf("bestMove: %v", bm)
		if bm.Target != 0 && bm.Target != 2 && bm.Target != 6 && bm.Target != 8 {
			t.Error(bm)
		}
	}
}

func Test_Minimax4(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PE, PE, PE,
		PE, PE, PE,
		PE, PE, PE}
	b.isXTurn = true
	stats := minimax.NewStats()
	minimax.Minimax(b, stats, 9)
	t.Logf("nNodes: %v", stats.NNodes)
	//bm, _ := stats.PosTable[b.Hash()].BestMove.(Move)
}
