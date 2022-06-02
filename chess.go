package chess

import (
	"errors"
	"fmt"
)

// Piece represents a piece on a chessboard, or "empty"
type Piece uint8

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

// abs gets the absolute value of an 8-bit integer.
func abs(x int8) int8 {
	var sb = x >> 7
	return (x ^ sb) + (sb & 1)
}

// Move represents a chess move on a board, with a piece, and source and destination squares.
// For a castle, move a king onto a rook as if you are taking it.
type Move struct {
	piece   Piece
	srcCol  uint8
	srcRow  uint8
	destCol uint8
	destRow uint8
	castle  bool
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
		if srcCol == destCol && 0 < destRow-srcRow <= 2 {
			valid = true
			break
		}
		if destRow == srcRow+1 && abs(destCol-srcCol) == 1 {
			valid = true
			break
		}
	case BlackPawn:
		if srcCol == destCol && 0 < srcRow-destRow <= 2 {
			valid = true
			break
		}
		if destRow == srcRow-1 && abs(destCol-srcCol) == 1 {
			valid = true
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
		piece:   p,
		srcCol:  srcCol,
		srcRow:  srcRow,
		destCol: destCol,
		destRow: destRow,
		castle:  castle,
	}, nil
}

// IsLegal determines whether a move is legal to do given a certain game context.
func (m *Move) IsLegal(g *Game) bool {

	if g.Board[m.srcRow][m.srcCol] != m.piece {
		return false // we can't move a piece that doesn't exist
	}
	// TODO: A nice way to check move legality

	return true
}

// Game represents a game of chess.
type Game struct {
	Board          [8][8]Piece
	MoveCounter    int32 // ??
	BlackCanCastle bool
	WhiteCanCastle bool
}

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

// NewGame initializes a new Game object
func NewGame() *Game {
	g := Game{
		Board: {{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
			{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
			{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook}},
		WhitePlayer:    white,
		BlackPlayer:    black,
		BlackCanCastle: true,
		WhiteCanCastle: true,
	}
	return &g
}

// DoMove executes a move on the board.
func (g *Game) DoMove(m *Move) error {
	// does the piece exist?
	if g.board[srcRow][srcCol] != m.piece {
		return errors.New("Game.Move: no such piece on that source square")
	}

	g.board[srcRow][srcCol] = Empty
	g.board[destRow][destCol] = m.piece
}

// DetermineCheck determines whether any players are in check.
func (g *Game) DetermineCheck() error {
	// TODO write the function
}
