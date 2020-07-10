package minimax

import (
	"fmt"
	"log"
	"math"
	"os"
)

type ZeroSumGame interface {
	ZCalcLegalMoves() []Move
	ZMakeMove(Move) bool // return isValid
	// undo the last ZMakeMove
	TakeBack() // undo the last Move
	// Evaluation can only be -1, 0 or 1 at the end of a game,
	// but usually this func is a estimation value of a position (heuristic).
	// :returns[0] bool: isExact score (game over) or not
	Evaluate() (bool, float64)
	// check whether if current turn is player who prefer max evaluation
	IsMaxPlayerTurn() bool
	// Hash must be unique for a position,
	// Forsyth-Edwards Notation can be used for human readable
	Hash() string
}

// Move examples: Chess{source: e2, target: e4}, TicTacToe{target: 4}
type Move interface {
	// CheckEqual usually is "=="
	CheckEqual(Move) bool
}

// Minimax chessprogramming.org/Minimax,
// :params PosTable: are passed and modified by all recursion steps,
// :params board: must be unchanged after recursion steps call MakeMove and TakeBack
func Minimax(board ZeroSumGame, stats *Stats, depth int) float64 {
	defer func() {
		stats.NNodes += 1
	}()
	posTable := stats.PosTable
	hash := board.Hash()

	var goodMove Move // best move in a shallow search
	if pos, found := posTable[hash]; found {
		if pos.IsTheEnd {
			return pos.Score
		}
		if depth <= pos.Depth {
			return pos.Score
		}
		goodMove = pos.BestMove
	}

	isTheEnd, score := board.Evaluate()

	if isTheEnd {
		posTable[hash] = Transposition{IsTheEnd: true, Depth: depth, Score: score}
		return score
	}

	if depth == 0 {
		posTable[hash] = Transposition{Depth: 0, Score: score}
		return score
	}

	moves := board.ZCalcLegalMoves()
	debug("hash: %v, depth: %v, moves: %v", hash, depth, moves)
	if len(moves) == 0 {
		posTable[hash] = Transposition{IsTheEnd: true, Score: score}
		return score
	}

	// reorder moves to get a better branch cut in Alpha-Beta,
	// does not help in this Minimax though
	if goodMove != nil {
		for i, move := range moves {
			if move.CheckEqual(goodMove) {
				moves[0], moves[i] = moves[i], moves[0]
				break
			}
		}
	}

	bestScore := math.Inf(1)
	if board.IsMaxPlayerTurn() {
		bestScore = math.Inf(-1)
	}
	bestMove := moves[0]
	for _, move := range moves {
		debug("hash %v about to go child %v", board.Hash(), move)
		board.ZMakeMove(move)
		childScore := Minimax(board, stats, depth-1)
		debug("hashAfterChild %v: %v, score: %v, oldBest: %v",
			move, board.Hash(), childScore, bestScore)
		board.TakeBack()
		if board.IsMaxPlayerTurn() {
			if childScore > bestScore {
				bestScore = childScore
				bestMove = move
			}
		} else {
			if childScore < bestScore {
				bestScore = childScore
				bestMove = move
			}
		}
	}
	posTable[hash] = Transposition{Score: bestScore, Depth: depth, BestMove: bestMove}
	debug("fored: hash: %v, bestScore %v, bestMove: %v",
		hash, bestScore, bestMove)
	if false { // very heavy debug code
		for k, v := range posTable {
			debug("__posTableRow %v: %#v", k, v)
		}
	}
	return bestScore

}

// TranspositionTable stores results of previously performed searches,
// key of this map is hash of the position,
type TranspositionTable map[string]Transposition

// Transposition store best move for a position,
// visit below url for more optimizing methods:
// https://www.chessprogramming.org/Transposition_Table#What_Information_is_Stored
type Transposition struct {
	IsTheEnd bool // game over, can determine win, loss or draw
	Score    float64
	Depth    int  // meaningless if (IsTheEnd == true)
	BestMove Move // meaningless if (IsTheEnd == true) or (Depth == 0)
}

func CalcBestMove(board ZeroSumGame, depth int) Transposition {
	stats := NewStats()
	Minimax(board, stats, depth)
	debug("NNodes: %v, nPoses: %v", stats.NNodes, len(stats.PosTable))
	bestMove := stats.PosTable[board.Hash()]
	if false { // debug zone
		log.Printf("NNodes: %v, nPoses: %v\n", stats.NNodes, len(stats.PosTable))
		for k, v := range stats.PosTable {
			if v.Depth >= depth-1 {
				log.Printf("__posTableRow %v: %#v", k, v)
			}
		}
	}
	return bestMove
}

var (
	isDebug = false
	std     = log.New(os.Stderr, "", log.Lshortfile)
)

func debug(format string, v ...interface{}) {
	if isDebug {
		std.Output(2, fmt.Sprintf(format, v...))
	}
}

// Stats stores results of all search steps,
// the most important field is TranspositionTable,
// others fields is used for measuring performance.
type Stats struct {
	PosTable TranspositionTable
	NNodes   int
}

func NewStats() *Stats {
	return &Stats{
		PosTable: make(map[string]Transposition),
		NNodes:   0,
	}
}
