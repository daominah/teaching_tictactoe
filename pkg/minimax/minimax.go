package minimax

import (
	"math"
)

type ZeroSumGame interface {
	ZCalcLegalMoves() []Move
	ZMakeMove(Move) bool // return isValid
	TakeBack()           // undo the last Move
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
	posHash := board.Hash()

	var goodMove Move // best move in a shallow search
	if pos, found := posTable[posHash]; found {
		if pos.IsTheEnd {
			return pos.Score
		}
		if depth <= pos.Depth {
			return pos.Score
		}
		goodMove = pos.BestMove
	}

	isTheEnd, score := board.Evaluate()

	if depth == 0 {
		posTable[posHash] = Transposition{IsTheEnd: isTheEnd, Depth: 0, Score: score}
		return score
	}

	allMoves := board.ZCalcLegalMoves()
	if len(allMoves) == 0 {
		posTable[posHash] = Transposition{IsTheEnd: true, Score: score}
		return score
	}

	// reorder allMoves to get a better branch cut in Alpha-Beta,
	// does not help in this NegaMax though
	if goodMove != nil {
		for i, move := range allMoves {
			if move.CheckEqual(goodMove) {
				allMoves[0], allMoves[i] = allMoves[i], allMoves[0]
				break
			}
		}
	}

	max := -math.Inf(1)
	maxMove := allMoves[0]
	for _, move := range allMoves {
		board.ZMakeMove(move)
		childScore := -NegaMax(board, posTable, depth-1)
		board.TakeBack()
		if childScore > max {
			max = childScore
			maxMove = move
		}
	}
	posTable[posHash] = Transposition{Score: max, Depth: depth, BestMove: maxMove}
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
