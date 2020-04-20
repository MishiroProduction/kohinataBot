package core

/**
core/messageファイル
主にアウトプット側の汎用の処理を記述します。
*/

import (
	"github.com/bwmarrin/discordgo"
)

// TalkToText is post text message to discord channel.
func TalkToText(session *discordgo.Session, channel *discordgo.Channel, message string) error {
	_, err := session.ChannelMessageSend(channel.ID, message)
	return err
}
