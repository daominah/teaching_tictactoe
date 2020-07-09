package tictactoe

import (
	"errors"
	"math/rand"
	"strings"
	"time"

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
	legalMoves := make([]Move, 0)
	for i := 0; i < WIDTH*HEIGHT; i++ {
		if b.squares[i] == PE {
			legalMoves = append(legalMoves, Move{target: i})
		}
	}
	return legalMoves
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
	if len(b.CalcLegalMoves()) == 0 {
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
		b.isXTurn = false
		return true
	}
	b.squares[m.target] = PO
	b.isXTurn = true
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

// :return : square index
func (b *Board) CalcBestMove() (int, error) {
	legals := b.CalcLegalMoves()
	if len(legals) == 0 {
		return 0, ErrNoLegalMoves
	}

	return 0, ErrNotImplemented
}

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
	return b.String()
}

func (m Move) CheckEqual(minimaxMove minimax.Move) bool {
	tttMove, ok := minimaxMove.(Move)
	if !ok {
		return false
	}
	return m.target == tttMove.target
}
