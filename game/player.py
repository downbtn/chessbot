class Player():
    def __init__(self, user: discord.User, elo: int = 800):
        self.id = user.id
        self.elo = elo
