package tictactoe

import (
	"math"
	"testing"

	"github.com/daominah/teaching_tictactoe/pkg/minimax"
)

func Test_AlphaBeta(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PX, PX, PE,
		PX, PO, PE,
		PO, PE, PO}
	b.isXTurn = true

	stats0 := minimax.NewStats()
	minimax.Minimax(b, stats0, 9)
	t.Logf("minimax nNode: %v, nPos: %v, bestMove: %v",
		stats0.NNodes, len(stats0.PosTable), stats0.PosTable[b.Hash()].BestMove)

	stats := minimax.NewStats()
	minimax.IsDebug = true
	minimax.AlphaBeta(b, stats, 9, -math.Inf(1), math.Inf(1))
	minimax.IsDebug = false
	t.Logf("nNodes: %v", stats.NNodes)
	bm, _ := stats.PosTable[b.Hash()].BestMove.(Move)
	t.Logf("bestMove: %#v", bm)
	if bm.Target != 0 && bm.Target != 2 && bm.Target != 6 && bm.Target != 8 {
		t.Error(bm)
	}
}

func Test_AlphaBeta2(t *testing.T) {
	b := NewBoard()
	b.squares = []Piece{
		PE, PE, PE,
		PE, PX, PE,
		PE, PE, PE}
	b.isXTurn = false

	for i := 0; i < 5; i++ {
		stats := minimax.NewStats()
		minimax.AlphaBeta(b, stats, 999, math.Inf(-1), math.Inf(1))
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
