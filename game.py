#!usr/bin/env python3
import chess
import discord
from PIL import Image

IMG_TEMPLATE = "./assets/board.png"
IMG_BN = "./assets/bN.png"
IMG_BB = "./assets/bB.png"
IMG_BK = "./assets/bK.png"
IMG_BQ = "./assets/bQ.png"
IMG_BR = "./assets/bR.png"
IMG_BP = "./assets/bP.png"
IMG_WN = "./assets/wN.png"
IMG_WB = "./assets/wB.png"
IMG_WK = "./assets/wK.png"
IMG_WQ = "./assets/wQ.png"
IMG_WR = "./assets/wR.png"
IMG_WP = "./assets/wP.png"

class Player():
    def __init__(self, user: discord.User, elo: int = 800):
        self.id = user.id
        self.elo = elo
        # Game history should be added later, but that seems complicated until
        # I can figure out databases


"""Game represents a running game of chess, including things like board state,
move history and players"""
class Game():
    def __init__(self, white_player: Player, black_player: Player):
        """Creates Game object and also initializes board state"""
        # Should this behavior be changed to have a separate init function? I have NO IDEA
        self.p_white = white_player
        self.p_black = black_player
        self.board = chess.Board()

    def render(board, of: str) -> None:
        # TODO: verify of is valid path
        with Image.open(IMG_TEMPLATE) as img:
            img.load()
            for row in range(8):
                for col in range(8):
                    # row 0 = 8th rank
                    # col 0 = a file
                    current_piece = self.board.piece_at(chess.square(col, 7 - row))
                    # wish there were a cleaner way to do this. lol
                    if current_piece == None:
                        continue
                    if current_piece.piece_type == chess.KNIGHT:
                        # Check for knight
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BN) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WN) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                    elif current_piece.piece_type == chess.KING:
                        # Check for king
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BK) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WK) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                    elif current_piece.piece_type == chess.BISHOP:
                        # Check for bishop
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BB) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WB) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                    elif current_piece.piece_type == chess.ROOK:
                        # Check for rook
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BR) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WR) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                    elif current_piece.piece_type == chess.QUEEN:
                        # Check for queen
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BQ) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WQ) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                    elif current_piece.piece_type == chess.PAWN:
                        # Check for pawn
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BP) as sprite:
                                img.paste(sprite, (90*row, 90*col))
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WP) as sprite:
                                img.paste(sprite, (90*row, 90*col))


    def move(self, move: chess.Move):
        self.board.push(move)