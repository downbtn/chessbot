package chess

import (
	"errors"
	"fmt"
)

// Piece represents a piece on a chessboard, or "empty"
type Piece uint8

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
	colorBlack uint8 = 0
	colorWhite       = 1
	colorEmpty       = 2
)

/* const NewGamePosition [8][8]Piece = {{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook}}
*/

// GetColor returns the color of a piece - 0 for black, 1 for white, 2 for empty
func (p Piece) getColor() uint8 {
	if p == Empty {
		return colorEmpty
	} else if 0 < p && p <= 6 {
		return colorWhite
	} else {
		return colorBlack
	}
}

// abs gets the absolute value of an 8-bit integer.
func abs(x int8) int8 {
	var sb = x >> 7
	return (x ^ sb) + (sb & 1)
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

// DetermineCheck determines whether a position on a board is in check or not.
// Returns a color value of the color in check.
func (b *Board) DetermineCheck() uint8 {
	// TODO
	return 0
}

// IsLegal determines whether a move is legal to do given a certain game context.
func (m *Move) IsLegal(b *Board) bool {
	playerColor := m.piece.getColor()
	if b[m.srcRow][m.srcCol] != m.piece {
		return false // we can't move a piece that doesn't exist
	}
	if b[m.destRow][m.destCol].getColor() == m.piece.getColor() && (!m.castle) {
		return false // we can't take our own piece unless the move is representing a castle
	}
	if m.pawnTaking && b[m.destRow][m.destCol].getColor() != (m.piece.getColor()^1) {
		return false // there is no piece of the other color to take
		// TODO: en passant doesn't work yet :P
	}

	// After all the other checks determine if the position after the move is a check
	newBoard := *b
	newBoard.doMove(m)
	if newBoard.DetermineCheck() == playerColor {
		return false // we can't put ourself in check!
	}
	return true
}

func (b *Board) doMove(m *Move) {
	b[m.srcRow][m.srcCol] = Empty
	b[m.destRow][m.destCol] = m.piece
}
