package main

/**
Mainファイル
このファイルは基本的に編集しない
*/

import (
	"fmt"
	"log"
	"os"

	"kohinataBot/core"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	// BotClientID is Client ID of this Bot
	BotClientID string
	// BotRoutine is channel for goroutine
	BotRoutine = make(chan bool)
	// VoiceSession is Session of Voice Connection
	VoiceSession *discordgo.VoiceConnection
)

func main() {
	// Read Environment Strings from dotenv.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error: Load dotenv failure")
	}
	// Initializing DiscordGo.
	discord, err := discordgo.New()
	if err != nil {
		log.Fatal("Error: Initializing DiscordGo failure")
	}
	// Set Discord Token and BotClientID variable.
	discord.Token = fmt.Sprintf("Bot %s", os.Getenv("BOT_TOKEN"))
	BotClientID = fmt.Sprintf("<@%s>", os.Getenv("BOT_CLIENT_ID"))
	// Set Handler.
	discord.AddHandler(core.CoreHandler)
	// Open WebSocket.
	if err := discord.Open(); if err != nil {
		log.Fatal("Error: Open Discord WebSocket failure")
	}
	<-BotRoutine
	return
}
