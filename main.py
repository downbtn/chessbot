#!/usr/bin/env python3
import discord
from discord import app_commands
from datetime import timedelta, timezone, datetime

import secret

intents = discord.Intents.default()
intents.message_content = True
client = discord.Client(intents=intents)
tree = app_commands.CommandTree(client)


"""Ping the bot and get latency"""
@tree.command(name="ping", description="Ping the bot", guilds=secret.guilds)
async def ping(interaction) -> None:
    old_timestamp = interaction.created_at
    new_timestamp = datetime.now(timezone.utc)
    diff_ms = int((new_timestamp - old_timestamp) / timedelta(milliseconds=1))
    await interaction.response.send_message(f"pong {diff_ms}ms :ping_pong:")


@client.event
async def on_ready():
    # Sync slash commands to whitelisted guilds
    for guild in secret.guilds:
        await tree.sync(guild=guild)
    print("Logged in as", client.user)


client.run(secret.token)
