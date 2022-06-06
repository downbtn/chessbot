package chess

import "errors"

// CastleAbility represents the ability of players to castle either queenside or kingside in a game.
type CastleAbility struct {
	BlackQ bool
	BlackK bool
	WhiteQ bool
	WhiteK bool
}

// Game represents a game of chess.
type Game struct {
	Board        Board
	MoveCounter  int32 // ??
	WhoCanCastle CastleAbility
	WhoseTurn    Color
}

// NewGame initializes a new Game object
func NewGame() *Game {
	ca := CastleAbility{BlackK: true,
		BlackQ: true,
		WhiteK: true,
		whiteQ: true,
	}
	g := Game{
		Board: Board{{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
			{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty},
			{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
			{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook}},
		MoveCounter:  0,
		WhoCanCastle: ca,
		WhoseTurn:    colorWhite,
	}
	return &g
}

func (g *Game) Turn(m *Move) error {
	if m.piece.getColor() != g.WhoseTurn {
		return errors.New("Game.Turn: not your turn")
	}
	if !m.IsLegal(g.board) {
		return errors.New("Game.Turn: illegal move")
	}
	if m.piece.getType() == Rook {
		// update castle ability!
	}
}

// IsLegal determines whether a move is legal to do given a certain game context.
func (g *Game) IsLegal(m *Move) bool {
	playerColor := m.piece.getColor()
	if g.Board[m.srcRow][m.srcCol] != m.piece {
		return false // we can't move a piece that doesn't exist
	}
	if m.castle {
		// are you able to castle?
		switch {
		case playerColor == colorBlack && m.destCol == 7:
			if !g.WhoCanCastle.BlackK {
				return false
			}
		case playerColor == colorBlack && m.destCol == 0:
			// black queenside
			if !g.WhoCanCastle.BlackQ {
				return false
			}

		case playerColor == colorWhite && m.destCol == 7:
			// white kingside
			if !g.WhoCanCastle.WhiteK {
				return false
			}

		case playerColor == colorWhite && m.destCol == 0:
			// white queenside
			if !g.WhoCanCastle.WhiteQ {
				return false
			}
		}
		// is anything blocking the castle?
		if m.srcCol < m.destCol {
			// iterate from srcCol up to destCol
			for col := 1; col < destCol-srcCol; col++ {
				if g.board[srcRow][srcCol+col] != Empty {
					return false
				}
			}
		} else {
			for col := -1; col > destCol-srcCol; col-- {
				if g.board[srcRow][srcCol+col] != Empty {
					return false
				}
			}
		}
	} else if g.Board[m.destRow][m.destCol].getColor() == playerColor {
		return false // we can't take our own piece unless the move is representing a castle
	}

	if m.pawnTaking && g.Board[m.destRow][m.destCol].getColor() != (m.piece.getColor()^1) {
		return false // there is no piece of the other color to take
		// TODO: en passant doesn't work yet :P
	}

	// After all the other checks determine if the position after the move is a check
	newBoard := g.Board
	newBoard.doMove(m)
	if newBoard.DetermineCheck() == playerColor {
		return false // we can't put ourself in check!
	}
	return true
}
