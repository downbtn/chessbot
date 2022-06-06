package chess

import (
	"errors"
	"fmt"
)

// Piece represents a piece on a chessboard, or "empty"
type Piece uint8

// PieceType represents a type of piece (e.g. pawn, king)
type PieceType uint8

// CheckState represents the state of who is in check on a board.
type CheckState uint8

// Color represents a color
type Color uint8

// Board represents the state of a board
type Board [8][8]Piece

// Move represents a chess move on a board, with a piece, and source and destination squares.
// For a castle, move a king onto a rook as if you are taking it.
type Move struct {
	piece      Piece
	srcCol     int8
	srcRow     int8
	destCol    int8
	destRow    int8
	castle     bool
	pawnTaking bool
}

type coords struct {
	row int8
	col int8
}

// Here we define all the constants representing various pieces.
const (
	Empty       Piece = 0
	WhiteKing         = 1
	WhiteQueen        = 2
	WhitePawn         = 3
	WhiteBishop       = 4
	WhiteKnight       = 5
	WhiteRook         = 6
	BlackKing         = 7
	BlackQueen        = 8
	BlackPawn         = 9
	BlackBishop       = 10
	BlackKnight       = 11
	BlackRook         = 12
)

const (
	colorBlack Color = 0
	colorWhite       = 1
	colorEmpty       = 2
)

const (
	NoCheck      CheckState = 0
	WhiteChecked            = 1
	BlackChecked            = 2
	WhiteMated              = 3
	BlackMated              = 4
)

const (
	EmptyType PieceType = 0
	King                = 1
	Queen               = 2
	Pawn                = 3
	Bishop              = 4
	Knight              = 5
	Rook                = 6
)

var NewGamePosition = Board{{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
	{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
	{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
	{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
	{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
	{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
	{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
	{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook}}

// abs gets the absolute value of an 8-bit integer.
func abs(x int8) int8 {
	var sb = x >> 7
	return (x ^ sb) + (sb & 1)
}

// GetColor returns the color of a piece - 0 for black, 1 for white, 2 for empty
func (p Piece) getColor() Color {
	if p == Empty {
		return colorEmpty
	} else if 0 < p && p <= 6 {
		return colorWhite
	} else {
		return colorBlack
	}
}

func (p Piece) getType() PieceType {
	if p == Empty {
		return EmptyType
	}
	return ((p - 1) % 6) + 1
}

// NewMove creates a new Move object describing a valid chess move.
func NewMove(p Piece, srcCol int8, srcRow int8, destCol int8, destRow int8) (*Move, error) {
	// Check that the move is valid
	// Note that we do not check that the move is *legal* - that requires a board position

	if srcCol < 0 || srcCol > 7 || srcRow < 0 || srcRow > 7 {
		return nil, fmt.Errorf("NewMove: source square is off the board row:%d column:%d", srcRow, srcCol)
	}
	if destCol < 0 || destCol > 7 || destRow < 0 || destRow > 7 {
		return nil, fmt.Errorf("NewMove: destination square is off the board row:%d column:%d", destRow, destCol)
	}
	if srcCol == destCol && srcRow == destRow {
		return nil, errors.New("NewMove: you can't move nowhere")
	}
	if p == Empty {
		return nil, errors.New("NewMove: you can't move an empty piece")
	}

	// can the piece actually go there?
	valid := false
	castle := false
	pawnTaking := false
	switch p {
	case WhiteKing:
		if srcCol == 4 && srcRow == 0 && destRow == 0 && (destCol == 0 || destCol == 7) {
			valid = true
			castle = true
			break
		}
		fallthrough
	case BlackKing:
		// First check castle!
		if srcCol == 4 && srcRow == 7 && destRow == 7 && (destCol == 0 || destCol == 7) {
			valid = true
			castle = true
			break
		}
		vDist := abs(srcRow - destRow)
		hDist := abs(srcCol - destCol)
		if hDist <= 1 && vDist <= 1 {
			valid = true
			break
		}

	case WhiteQueen:
		fallthrough
	case BlackQueen:
		// Same row or column?
		if destRow == srcRow || destCol == srcCol {
			valid = true
			break
		}
		// Diagonal?
		if abs(srcRow-destRow) == abs(srcCol-destCol) {
			valid = true
			break
		}

	case WhitePawn:
		if srcCol == destCol && 0 < destRow-srcRow && destRow-srcRow <= 2 {
			valid = true
			break
		}
		if destRow == srcRow+1 && abs(destCol-srcCol) == 1 {
			valid = true
			pawnTaking = true
			break
		}
	case BlackPawn:
		if srcCol == destCol && 0 < srcRow-destRow && srcRow-destRow <= 2 {
			valid = true
			break
		}
		if destRow == srcRow-1 && abs(destCol-srcCol) == 1 {
			valid = true
			pawnTaking = true
			break
		}

	case WhiteBishop:
		fallthrough
	case BlackBishop:
		if abs(srcRow-destRow) == abs(srcCol-destCol) {
			valid = true
			break
		}

	case WhiteKnight:
		fallthrough
	case BlackKnight:
		vDist := abs(srcRow - destRow)
		hDist := abs(srcCol - destCol)
		if (vDist == 2 && hDist == 1) || (vDist == 1 && hDist == 2) {
			valid = true
			break
		}

	case WhiteRook:
		fallthrough
	case BlackRook:
		if destRow == srcRow || destCol == srcCol {
			valid = true
			break
		}
	}
	if !valid {
		return nil, errors.New("NewMove: the piece is not allowed to move there")
	}
	return &Move{
		piece:      p,
		srcCol:     srcCol,
		srcRow:     srcRow,
		destCol:    destCol,
		destRow:    destRow,
		castle:     castle,
		pawnTaking: pawnTaking,
	}, nil
}

func (b *Board) threatens(src *coords, dst *coords) bool {
	p := b[src.row][src.col]
	switch p {
	case WhiteKing:
		fallthrough
	case BlackKing:
		vDist := abs(src.row - dst.row)
		hDist := abs(src.col - dst.col)
		if hDist <= 1 && vDist <= 1 {
			return true
		}

	case WhiteQueen:
		fallthrough
	case BlackQueen:
		// Same row or column?
		if dst.row == src.row {
			// same row - is it blocked?
			if dst.col < src.col {
				for i := dst.col + 1; i < src.col; i++ {
					if b[dst.row][i] != Empty {
						return false
					}
				}
			} else {
				for i := src.col + 1; i < dst.col; i++ {
					if b[dst.row][i] != Empty {
						return false
					}
				}
			}
			return true
		}
		if dst.col == src.col {
			if dst.row < src.row {
				for i := dst.row + 1; i < src.row; i++ {
					if b[i][dst.col] != Empty {
						return false
					}
				}
			} else {
				for i := src.row + 1; i < dst.row; i++ {
					if b[i][dst.col] != Empty {
						return false
					}
				}
			}
			return true
		}
		// Diagonal?
		vDiff := src.row - dst.row
		hDiff := src.col - dst.col
		switch {
		case vDiff > 0 && vDiff == hDiff:
			// i.e. src.row > dst.row src.col > dst.col
			for i := 1; i < vDiff; i++ {
				if b[dst.row+i][dst.col+i] != Empty {
					return false
				}
			}
			return true
		case vDiff < 0 && vDiff == hDiff:
			// i.e. src.row < dst.row src.col < dst.col
			for i := -1; i > vDiff; i-- {
				if b[dst.row+i][dst.col+i] != Empty {
					return false
				}
			}
			return true
		case vDiff > 0 && vDiff == -hDiff:
			// src.row > dst.row src.col < dst.col
			for i := 1; i < vDiff; i++ {
				if b[dst.row+i][dst.col-i] != Empty {
					return false
				}
			}
			return true
		case vDiff < 0 && vDiff == -hDiff:
			// i.e. src.row < dst.row src.col > dst.col
			for i := -1; i > vDiff; i-- {
				if b[dst.row-i][dst.col+i] != Empty {
					return false
				}
			}
			return true
		}

	case WhitePawn:
		if dst.row == src.row+1 && abs(dst.col-src.col) == 1 {
			return true
		}
	case BlackPawn:
		if dst.row == src.row-1 && abs(dst.col-src.col) == 1 {
			return true
		}

	case WhiteBishop:
		fallthrough
	case BlackBishop:
		vDiff := src.row - dst.row
		hDiff := src.col - dst.col
		switch {
		case vDiff > 0 && vDiff == hDiff:
			// i.e. src.row > dst.row src.col > dst.col
			for i := 1; i < vDiff; i++ {
				if b[dst.row+i][dst.col+i] != Empty {
					return false
				}
			}
			return true
		case vDiff < 0 && vDiff == hDiff:
			// i.e. src.row < dst.row src.col < dst.col
			for i := -1; i > vDiff; i-- {
				if b[dst.row+i][dst.col+i] != Empty {
					return false
				}
			}
			return true
		case vDiff > 0 && vDiff == -hDiff:
			// src.row > dst.row src.col < dst.col
			for i := 1; i < vDiff; i++ {
				if b[dst.row+i][dst.col-i] != Empty {
					return false
				}
			}
			return true
		case vDiff < 0 && vDiff == -hDiff:
			// i.e. src.row < dst.row src.col > dst.col
			for i := -1; i > vDiff; i-- {
				if b[dst.row-i][dst.col+i] != Empty {
					return false
				}
			}
			return true

		}

	case WhiteKnight:
		fallthrough
	case BlackKnight:
		vDist := abs(src.row - dst.row)
		hDist := abs(src.col - dst.col)
		if (vDist == 2 && hDist == 1) || (vDist == 1 && hDist == 2) {
			return true
		}

	case WhiteRook:
		fallthrough
	case BlackRook:
		if dst.row == src.row {
			// same row - is it blocked?
			if dst.col < src.col {
				for i := dst.col + 1; i < src.col; i++ {
					if b[dst.row][i] != Empty {
						break
					}
				}
			} else {
				for i := src.col + 1; i < dst.col; i++ {
					if b[dst.row][i] != Empty {
						break
					}
				}
			}
			return true
		}
		if dst.col == src.col {
			if dst.row < src.row {
				for i := dst.row + 1; i < src.row; i++ {
					if b[i][dst.col] != Empty {
						break
					}
				}
			} else {
				for i := src.row + 1; i < dst.row; i++ {
					if b[i][dst.col] != Empty {
						break
					}
				}
			}
			return true
		}
	}
	return false
}

// DetermineCheck determines whether a position on a board is in check or not.
// Returns a color value of the color in check.
func (b *Board) DetermineCheck() CheckState {
	// there has to be a better way to do this
	var wKingSquare coords
	var bKingSquare coords

	whiteChecked := false
	blackChecked := false

	// this seems kinda unnecessary - maybe store king coords in game object?
	for i, row := range b {
		for j, square := range row {
			if square == WhiteKing {
				wKingSquare = coords{col: j, row: i}
			}
			if square == BlackKing {
				bKingSquare = coords{col: j, row: i}
			}
		}
	}

	for i, row := range b {
		for j, piece := range row {
			if piece == Empty || piece.getType() == King {
				continue
			}
			pCoords := coords{col: j, row: i}
			if piece.getColor() == colorBlack {
				if b.threatens(pCoords, wKingSquare) {
					whiteChecked = true
				}
			} else {
				if b.threatens(pCoords, bKingSquare) {
					blackChecked = true
				}
			}
		}
	}

	if whiteChecked {
		return WhiteChecked
	}
	if blackChecked {
		return BlackChecked
	}
	return NoCheck
}

func (b *Board) doMove(m *Move) {
	if m.castle {
		// queenside i.e. destCol == 0 or kingside i.e. destCol == 7?
		if m.destCol == 0 {
			// castle queenside
			// king col 4 -> 2
			// rook col 0 -> 3
			b[m.srcRow][3] = b[m.destRow][m.destCol]
			b[m.srcRow][4] = Empty
			b[m.srcRow][0] = Empty
			b[m.srcRow][2] = m.piece
		} else if m.destCol == 7 {
			// castle kingside
			// king col 4 -> 6
			// rook col 7 -> 5
			b[m.srcRow][5] = b[m.destRow][m.destCol]
			b[m.srcRow][6] = m.piece
			b[m.srcRow][4] = Empty
			b[m.srcRow][7] = Empty
		}
	} else {
		b[m.srcRow][m.srcCol] = Empty
		b[m.destRow][m.destCol] = m.piece
	}
}
