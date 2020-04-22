package core

/**
core/coreファイル
Botのコアとなる動作を記述していきます。
*/

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// CoreHandler is handler of bot
func CoreHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		log.Println("Get ChannelID failure: ", err.Error())
		return
	}
	// 発信元がBotの場合以降の処理を行わない
	if message.Author.Bot {
		return
	}
	// 基本的なBotの処理
	switch {
	case strings.Contains(message.Content, "おはよう"):
		err = TalkToText(session, channel, "おはようございます、プロデューサーさん！")

	case strings.Contains(message.Content, "こんにちは"):
		err = TalkToText(session, channel, "いい天気ですね、プロデューサーさん♪")

	case strings.Contains(message.Content, "こんばんは"):
		err = TalkToText(session, channel, "いい夜ですね、プロデューサーさん♪")

	case strings.Contains(message.Content, "おやすみなさい"):
		err = TalkToText(session, channel, "おやすみなさい、プロデューサーさん…")
	}
	if err != nil {
		log.Println("Post Message failure: ", err.Error())
	}
	return
}
