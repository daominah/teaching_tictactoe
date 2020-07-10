package minimax

// AlphaBeta
// :params PosTable: are passed and modified by all recursion steps,
// :params board: must be unchanged after recursion steps call MakeMove and TakeBack
// :params β: help to break while finding max,
//   assume we are calling this func on a node (named CNode) that IsMaxPlayerTurn,
//   and had a score of CNode's sibling (passed as β). While looping through
//   CNode's children, if a child return score >= β, we can stop
//   the loop, and confirmed CNode's score < alpha.
func AlphaBeta(board ZeroSumGame, stats *Stats, depth int,
	alpha float64, beta float64) float64 {
	defer func() {
		stats.NNodes += 1
	}()
	posTable := stats.PosTable
	hash0 := board.Hash()

	// α, β were received from parent,
	// can be improved by TranspositionTable if we searched the position,
	// at the beginning of the root, we call with {alpha: -inf, beta: +inf}
	thisAlpha := alpha
	thisBeta := beta
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
		} else { // score is just a bound, not exact
			if board.IsMaxPlayerTurn() {
				// TODO
			} else {
				// TODO
			}
		}
	}

	isTheEnd, score := board.Evaluate()

	if isTheEnd {
		posTable[hash0] = Transposition{IsTheEnd: true, Depth: depth, Score: score}
		return score
	}

	if depth == 0 {
		posTable[hash0] = Transposition{Depth: 0, Score: score}
		return score
	}

	moves := board.ZCalcLegalMoves()
	debug("hash0: %v, depth: %v, moves: %v", hash0, depth, moves)
	if len(moves) == 0 {
		posTable[hash0] = Transposition{IsTheEnd: true, Depth: depth, Score: score}
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
	isCutNode := false
	boundMove := moves[0] // if !isCutNode: boundMove is bestMove
	for i, move := range moves {
		board.ZMakeMove(move)
		debug("go child %v of %v: %v",
			move, !board.IsMaxPlayerTurn(), board.Hash())
		childScore := AlphaBeta(board, stats, depth-1, thisAlpha, thisBeta)
		debug("score go child %v of %v: %v, score: %v",
			move, !board.IsMaxPlayerTurn(), board.Hash(), childScore)
		board.TakeBack()

		if board.IsMaxPlayerTurn() {
			if childScore > thisAlpha { // alpha = max(childScore, alpha)
				thisAlpha = childScore
				boundMove = move
				if thisAlpha >= beta { // skip remaining children, go to return
					debug("α>β cut: hash0: %v child: %v, tα: %v, β: %v",
						hash0, move, thisAlpha, beta)
					if i != len(moves)-1 {
						isCutNode = true
					}
					break
				}
			}
		} else {
			if childScore < beta { // beta = min(childScore, beta)
				thisBeta = childScore
				boundMove = move
				if thisBeta <= alpha {
					debug("β<α cut: hash0: %v child: %v, α: %v, tβ: %v",
						board.Hash(), move, alpha, thisBeta)
					if i != len(moves)-1 {
						isCutNode = true
					}
					break
				}
			}
		}
	}
	var boundScore float64
	if board.IsMaxPlayerTurn() {
		boundScore = thisAlpha
	} else {
		boundScore = thisBeta
	}
	posTable[hash0] = Transposition{
		IsCutNode: isCutNode, Score: boundScore, BestMove: boundMove,
		Depth: depth}
	debug("fored: hash0: %v, boundScore %v, boundMove: %v",
		hash0, boundScore, boundMove)
	if true { // very heavy debug code
		for k, v := range posTable {
			debug("__posTableRow %v: %#v", k, v)
		}
	}
	return boundScore
}
