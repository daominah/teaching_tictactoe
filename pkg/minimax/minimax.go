package minimax

import (
	"log"
	"math"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type ZeroSumGame interface {
	ZCalcLegalMoves() []Move
	ZMakeMove(Move) bool // return isValid
	// undo the last ZMakeMove
	TakeBack() // undo the last Move
	// Evaluation can only be -1, 0 or 1 at the end of a game,
	// but usually this func is a estimation value of a position (heuristic).
	// :returns[0] bool: isExact score (game over) or not
	// For NegaMax to work, this func must return a score relative
	// to the side to being evaluated.
	Evaluate() (bool, float64)
	// Hash must be unique for a position,
	// Forsyth-Edwards Notation can be used for human readable
	Hash() string
}

// Move examples: Chess{source: e2, target: e4}, TicTacToe{target: 4}
type Move interface {
	// CheckEqual usually is "=="
	CheckEqual(Move) bool
}

// NegaMax chessprogramming.org/Negamax,
// :params board, posTable: are passed and modified by all recursion steps,
func NegaMax(board ZeroSumGame, posTable TranspositionTable, depth int) float64 {
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
	log.Printf("hash: %v, depth: %v, moves: %v", hash, depth, moves)
	if len(moves) == 0 {
		posTable[hash] = Transposition{IsTheEnd: true, Score: score}
		return score
	}

	// reorder moves to get a better branch cut in Alpha-Beta,
	// does not help in this NegaMax though
	if goodMove != nil {
		for i, move := range moves {
			if move.CheckEqual(goodMove) {
				moves[0], moves[i] = moves[i], moves[0]
				break
			}
		}
	}

	max := -math.Inf(1)
	maxMove := moves[0]
	for _, move := range moves {
		log.Printf("hash %v about to go child %v", board.Hash(), move)
		board.ZMakeMove(move)
		childScore := -NegaMax(board, posTable, depth-1)
		log.Printf(" hashAfterChild %v: %v, score: %v, oldMax: %v",
			move, board.Hash(), childScore, max)
		board.TakeBack()
		if childScore > max {
			max = childScore
			maxMove = move
		}
	}
	posTable[hash] = Transposition{Score: max, Depth: depth, BestMove: maxMove}
	log.Printf("fored: hash: %v, max %v, maxMove: %v", hash, max, maxMove)
	for k, v := range posTable {
		log.Printf("__posTableRow %v: %#v", k, v)
	}
	return max
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
	posTable := make(map[string]Transposition, int(math.Pow(2, float64(depth))))
	NegaMax(board, posTable, depth)
	bestMove := posTable[board.Hash()]
	return bestMove
}
