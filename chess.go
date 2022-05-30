package chess

import "errors"

type Piece uint8

const (
	empty       Piece = 0
	whiteKing         = 1
	whiteQueen        = 2
	whitePawn         = 3
	whiteBishop       = 4
	whiteKnight       = 5
	whiteRook         = 6
	blackKing         = 7
	blackQueen        = 8
	blackPawn         = 9
	blackBishop       = 10
	blackKnight       = 11
	blackRook         = 12
)

// Player represents a player in a chess game.
type Player struct {
	name    string
	country string // ISO-3166 alpha 2 country code
	rating  int
	id      int
}

// Game represents a game of chess.
type Game struct {
	board        [8][8]Piece
	whitePlayer  *Player
	blackPlayer  *Player
	whiteInCheck bool
	blackInCheck bool
	moveCounter  int32
}

// NewGame initializes a new Game object
func NewGame(Player *white, player *black) *Game {
	g := Game{
		board: {{whiteRook, whiteKnight, whiteBishop, whiteQueen, whiteKing, whiteBishop, whiteKnight, whiteRook},
			{whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn, whitePawn},
			{empty, empty, empty, empty, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty},
			{empty, empty, empty, empty, empty, empty, empty, empty},
			{blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn, blackPawn},
			{blackRook, blackKnight, blackBishop, blackQueen, blackKing, blackBishop, blackKnight, blackRook}},
		whiteInCheck: false,
		blackInCheck: false,
		whitePlayer:  white,
		blackPlayer:  black,
		moveCounter:  0,
	}
	return &g
}

// Move moves a piece on the board.
func (g *Game) Move(piece Piece, srcCol uint8, srcRow uint8, destCol uint8, destRow uint8) error {
	// check validity
	// does the piece exist?
	if piece == empty {
		return errors.New("Game.Move: you can't move an empty piece")
	}
	if g.board[srcRow][srcCol] != piece {
		return errors.New("Game.Move: no such piece on that source square")
	}
	// TODO: check if the position would put the mover in check

}

func (g *Game) DetermineCheck() error {
}
