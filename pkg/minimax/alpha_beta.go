package minimax

import (
	"math"
)

// AlphaBeta
// :params PosTable: are passed and modified by all recursion steps,
// :params board: must be unchanged after recursion steps call MakeMove and TakeBack
// :params β: help to break while finding max score of current node's children,
//   assume we are calling this func on a node (named CNode) that IsMaxPlayerTurn,
//   and had a score of CNode's sibling (passed as β). While looping through
//   CNode's children, if score(a child) >= β, so score(CNode) >= β, so CNode
//   cannot be min of CNode's parent.
func AlphaBeta(board ZeroSumGame, stats *Stats, depth int,
	alpha float64, beta float64) float64 {
	defer func() {
		stats.NNodes += 1
	}()
	posTable := stats.PosTable
	hash0 := board.Hash()

	var goodMove Move // best move in a shallow search
	if pos, found := posTable[hash0]; found {
		if pos.IsTheEnd {
			return pos.Score
		}
		if !pos.IsCutNode {
			if depth <= pos.Depth {
				return pos.Score
			}
			goodMove = pos.BestMove
		}
	}

	isTheEnd, score := board.Evaluate()

	if isTheEnd {
		posTable[hash0] = Transposition{IsTheEnd: true, Score: score}
		return score
	}

	if depth == 0 {
		posTable[hash0] = Transposition{Depth: 0, Score: score}
		return score
	}

	moves := board.ZCalcLegalMoves()
	debug("hash0: %v, depth: %v, moves: %v", hash0, depth, moves)
	if len(moves) == 0 {
		posTable[hash0] = Transposition{IsTheEnd: true, Score: score}
		return score
	}

	// moves order is very important,
	// user should return good moves first in func ZCalcLegalMoves
	if goodMove != nil {
		for i, move := range moves {
			if move.CheckEqual(goodMove) {
				moves[0], moves[i] = moves[i], moves[0]
				break
			}
		}
	}

	//
	isThisFindingMax := board.IsMaxPlayerTurn()
	isCutNode := false
	boundMove := moves[0]     // if !isCutNode: boundMove is bestMove
	boundScore := math.Inf(1) // best score among siblings
	if isThisFindingMax {
		boundScore = math.Inf(-1)
	}
	for i, move := range moves {
		board.ZMakeMove(move)
		debug("go child %v of %v: %v",
			move, !board.IsMaxPlayerTurn(), board.Hash())
		var childScore float64
		if isThisFindingMax { // so after ZMakeMove: child is finding min
			childScore = AlphaBeta(board, stats, depth-1, boundScore, +math.Inf(1))
		} else { // after ZMakeMove: child is finding max
			childScore = AlphaBeta(board, stats, depth-1, -math.Inf(1), boundScore)
		}
		debug("score go child %v of %v: %v, score: %v",
			move, !board.IsMaxPlayerTurn(), board.Hash(), childScore)
		board.TakeBack()

		if isThisFindingMax {
			if childScore > boundScore {
				boundScore = childScore
				boundMove = move
				if boundScore >= beta { // skip remaining children, go to return
					debug("β cut for %v on child: %v, score: %v, β: %v",
						hash0, move, boundScore, beta)
					if i != len(moves)-1 {
						isCutNode = true
					}
					break
				}
			}
		} else {
			if childScore < boundScore {
				boundScore = childScore
				boundMove = move
				if boundScore <= alpha {
					debug("α cut for %v on child: %v, score: %v, α: %v",
						hash0, move, boundScore, alpha)
					if i != len(moves)-1 {
						isCutNode = true
					}
					break
				}
			}
		}
	}

	if !isCutNode {
		posTable[hash0] = Transposition{IsCutNode: false,
			Score: boundScore, BestMove: boundMove, Depth: depth}
	} else {
		// TODO
	}
	debug("fored: hash0: %v, boundScore %v, boundMove: %v",
		hash0, boundScore, boundMove)
	if false { // very heavy debug code
		for k, v := range posTable {
			debug("__posTableRow %v: %#v", k, v)
		}
	}
	return boundScore
}
