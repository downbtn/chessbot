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

"""Game represents a running game of chess, including things like board state,
move history and players"""
class Game():
    def __init__(self, white_player: Player, black_player: Player):
        """Creates Game object and also initializes board state"""
        # Should this behavior be changed to have a separate init function? I have NO IDEA
        self.p_white = white_player
        self.p_black = black_player
        self.board = chess.Board()

    def render(self, of: str) -> None:
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
                        # Knight sprite dimensions are 62x70, therefore we need
                        # offset 14 on x and 10 on y to center
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BN) as sprite:
                                img.paste(sprite, (90*col + 14, 90*row + 10), sprite)
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WN) as sprite:
                                img.paste(sprite, (90*col + 14, 90*row + 10), sprite)
                    elif current_piece.piece_type == chess.KING:
                        # Check for king
                        # King sprite has dimensions 77x78
                        # Need offset 6, 6 to center
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BK) as sprite:
                                img.paste(sprite, (90*col + 6, 90*row + 6), sprite)
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WK) as sprite:
                                img.paste(sprite, (90*col + 6, 90*row + 6), sprite)
                    elif current_piece.piece_type == chess.BISHOP:
                        # Check for bishop
                        # Bishop sprite is 72x73
                        # Need offset 9, 9 to center
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BB) as sprite:
                                img.paste(sprite, (90*col + 9, 90*row + 9), sprite)
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WB) as sprite:
                                img.paste(sprite, (90*col + 9, 90*row + 9), sprite)
                    elif current_piece.piece_type == chess.ROOK:
                        # Check for rook
                        # Rook sprite is 60x66
                        # Need offset 15, 12 to center
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BR) as sprite:
                                img.paste(sprite, (90*col + 15, 90*row + 12), sprite)
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WR) as sprite:
                                img.paste(sprite, (90*col + 15, 90*row + 12), sprite)
                    elif current_piece.piece_type == chess.QUEEN:
                        # Check for queen
                        # Queen sprite is 77x70
                        # Need offset 7, 10 to center
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BQ) as sprite:
                                img.paste(sprite, (90*col + 7, 90*row + 10), sprite)
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WQ) as sprite:
                                img.paste(sprite, (90*col + 7, 90*row + 10), sprite)
                    elif current_piece.piece_type == chess.PAWN:
                        # Check for pawn
                        # Pawn sprite is 48x59
                        # Needs an offset of 21, 16 to center
                        if current_piece.color == chess.BLACK:
                            with Image.open(IMG_BP) as sprite:
                                img.paste(sprite, (90*col + 21, 90*row + 16), sprite)
                        elif current_piece.color == chess.WHITE:
                            with Image.open(IMG_WP) as sprite:
                                img.paste(sprite, (90*col + 21, 90*row + 16), sprite)

            img.save(of)



    def move(self, move: chess.Move):
        self.board.push(move)
