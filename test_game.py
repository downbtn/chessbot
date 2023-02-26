import unittest
import game
from discord import Object
import chess


class TestGameLogic(unittest.TestCase):
    def setUp(self):
        self.player = game.Player(Object(899049400867881002))
        self.game = game.Game(self.player, self.player)

    def test_player(self):
        self.assertEquals(self.player.elo, 800)
        self.assertEquals(self.player.id, 899049400867881002)

    def test_render(self):
        self.game.render("test_render.png")
        print("Manually evaluate rendered test_render.png plz")

    def test_move(self):
        self.game.move(chess.Move.from_uci("e2e4"))
        self.game.move(chess.Move.from_uci("e7e5"))
        self.game.move(chess.Move.from_uci("e1e2"))
        self.game.render("test_render_2.png")
        print("Manually evaluate rendered test_render_2.png plz")


if __name__ == "__main__":
    unittest.main()
