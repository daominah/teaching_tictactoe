package tictactoe

import (
	"strings"
)

type Board struct {
	isXTurn bool
	squares []Piece
}

type Piece string

const (
	PX Piece = "X"
	PO Piece = "O"
	PE Piece = "" // empty
)

const (
	WIDTH  = 3
	HEIGHT = 3
)

type Result string

const (
	Win     Result = "XWin"
	Draw    Result = "XDraw"
	Loss    Result = "XLoss"
	Playing Result = "Playing"
)

func NewBoard() *Board {
	b := &Board{
		isXTurn: true,
		squares: make([]Piece, WIDTH*HEIGHT),
	}
	for i := 0; i < WIDTH*HEIGHT; i++ {
		b.squares[i] = PE
	}
	return b
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

func (b *Board) CalcLegalMoves() []int {
	legalMoves := make([]int, 0)
	for i := 0; i < WIDTH*HEIGHT; i++ {
		if b.squares[i] == PE {
			legalMoves = append(legalMoves, i)
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
