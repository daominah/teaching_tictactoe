package tictactoe

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"fmt"

	"github.com/daominah/teaching_tictactoe/pkg/minimax"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Board struct {
	isXTurn bool
	squares []Piece
	history []Move
}

func NewBoard() *Board {
	b := &Board{
		isXTurn: true,
		squares: make([]Piece, WIDTH*HEIGHT),
		history: make([]Move, 0),
	}
	for i := 0; i < WIDTH*HEIGHT; i++ {
		b.squares[i] = PE
	}
	return b
}

type Move struct {
	// target square index
	target int
}

type Piece string

// enum Piece
const (
	PX Piece = "X"
	PO Piece = "O"
	PE Piece = "" // empty
)

// board size
const (
	WIDTH  = 3
	HEIGHT = 3
)

type Result string

// enum result of player who made the first move
const (
	Win     Result = "XWin"
	Draw    Result = "XDraw"
	Loss    Result = "XLoss"
	Playing Result = "Playing"
)

func (b *Board) CalcLegalMoves() []Move {
	if b.CheckResult() != Playing {
		return nil
	}
	legalMoves := make([]Move, 0)
	for i := 0; i < WIDTH*HEIGHT; i++ {
		if b.squares[i] == PE {
			legalMoves = append(legalMoves, Move{target: i})
		}
	}
	return legalMoves
}

func (b *Board) ZCalcLegalMoves() []minimax.Move {
	moves := b.CalcLegalMoves()
	ret := make([]minimax.Move, len(moves))
	for i, _ := range moves {
		ret[i] = moves[i]
	}
	return ret
}

func (b *Board) CheckResult() Result {
	lines := [][]int{
		[]int{0, 1, 2},
		[]int{0, 3, 6},
		[]int{0, 4, 8},
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{2, 4, 6},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	for _, piece := range []Piece{PX, PO} {
		for _, line := range lines {
			isSamePiece := true
			for _, sqrIdx := range line {
				if b.squares[sqrIdx] != piece {
					isSamePiece = false
					break
				}
			}
			if isSamePiece {
				if piece == PX {
					return Win
				} else {
					return Loss
				}
			}
		}
	}
	isFullBoard := true
	for _, p := range b.squares {
		if p == PE {
			isFullBoard = false
			break
		}
	}
	if isFullBoard {
		return Draw
	}
	return Playing
}

// :return : is valid move
func (b *Board) MakeMove(m Move) bool {
	if b.squares[m.target] != PE {
		return false
	}
	if b.isXTurn {
		b.squares[m.target] = PX
	} else {
		b.squares[m.target] = PO
	}
	b.isXTurn = !b.isXTurn
	b.history = append(b.history, m)
	return true
}

func (b *Board) TakeBack() {
	if len(b.history) < 1 {
		return
	}
	lastMove := b.history[len(b.history)-1]
	b.history = b.history[:len(b.history)-1]
	b.squares[lastMove.target] = PE
	b.isXTurn = !b.isXTurn
}

var (
	ErrNoLegalMoves   = errors.New("no legal moves are available")
	ErrNotImplemented = errors.New("not implemented")
)

// :return : square index
func (b *Board) CalcRandomMove() (Move, error) {
	legals := b.CalcLegalMoves()
	if len(legals) == 0 {
		return Move{}, ErrNoLegalMoves
	}
	bestMove := legals[rand.Intn(len(legals))]
	return bestMove, nil
}

// Evaluate for Negamax
func (b *Board) Evaluate() (bool, float64) {
	result := b.CheckResult()
	switch result {
	case Win:
		return true, 1
	case Loss:
		return true, -1
	case Draw:
		return true, 0
	default: // playing
		return false, 0
	}
}

func (b *Board) String() string {
	rowStrs := []string{""}
	for row := 0; row < HEIGHT; row++ {
		rowStr := make([]string, 0)
		for col := 0; col < WIDTH; col++ {
			sqrIdx := WIDTH*row + col
			pieceStr := string(b.squares[sqrIdx])
			if pieceStr == "" {
				pieceStr = "."
			}
			rowStr = append(rowStr, pieceStr)
		}
		rowStrs = append(rowStrs, strings.Join(rowStr, " "))
	}
	return strings.Join(rowStrs, "\n")
}

func (b *Board) Hash() string {
	oneLine := strings.ReplaceAll(b.String(), "\n", "|")
	turn := PX
	if !b.isXTurn {
		turn = PO
	}
	return fmt.Sprintf("%v|t%v", oneLine, turn)
}

func (b *Board) ZMakeMove(m minimax.Move) bool {
	tttMove, ok := m.(Move)
	if !ok {
		return false
	}
	return b.MakeMove(tttMove)
}

func (m Move) CheckEqual(minimaxMove minimax.Move) bool {
	tttMove, ok := minimaxMove.(Move)
	if !ok {
		return false
	}
	return m.target == tttMove.target
}

// :return : square index
func (b *Board) CalcBestMove() Move {
	moves := b.CalcLegalMoves()
	best := minimax.CalcBestMove(b, len(moves))
	bestMove := best.BestMove
	tttMove, _ := bestMove.(Move)
	return tttMove
}
