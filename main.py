#!/usr/bin/env python3
import discord
from discord import app_commands
from datetime import timedelta, timezone, datetime

import config

intents = discord.Intents.default()
intents.message_content = True
client = discord.Client(intents=intents)
tree = app_commands.CommandTree(client)


"""Ping the bot and get latency"""
@tree.command(name="ping", description="Ping the bot", guilds=config.guilds)
async def ping(interaction) -> None:
    old_timestamp = interaction.created_at
    new_timestamp = datetime.now(timezone.utc)
    diff_ms = int((new_timestamp - old_timestamp) / timedelta(milliseconds=1))
    embed = discord.Embed(title="Pong!",
                          description=f"{diff_ms}ms :ping_pong:",
                          color=discord.Colour.random())
    await interaction.response.send_message(embed=embed)


@client.event
async def on_ready():
    # Sync slash commands to whitelisted guilds
    for guild in config.guilds:
        await tree.sync(guild=guild)
    print("Logged in as", client.user)


client.run(config.TOKEN)
