package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages sent by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Handle the !poll command
	if strings.HasPrefix(m.Content, "!poll") {
		handlePollCommand(s, m)
	}
}

func handlePollCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.TrimPrefix(m.Content, "!poll ")
	args := strings.Split(content, ";")
	if len(args) < 3 {
		s.ChannelMessageSend(m.ChannelID, "Command format: !poll question; option1; option2; [option3; ...]")
		return
	}

	question := args[0]
	options := args[1:]

	// Create and send the poll message
	pollMessage, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("**%s**\n%s", question, formatOptions(options)))
	if err != nil {
		fmt.Println("Error sending poll message:", err)
		return
	}

	// Add reactions for voting
	for i := 0; i < len(options) && i < 9; i++ {
		emoji := fmt.Sprintf("%d⃣", i+1) // Use number emojis for options
		s.MessageReactionAdd(m.ChannelID, pollMessage.ID, emoji)
	}
}

// Format options for display in the message
func formatOptions(options []string) string {
	var builder strings.Builder
	for i, option := range options {
		builder.WriteString(fmt.Sprintf("%d. %s\n", i+1, option))
	}
	return builder.String()
}
