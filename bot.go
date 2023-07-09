package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var games []*TicTacToe

type Bot struct {
	token string
}

var prefix = "!"

func NewBot(token string) *Bot {
	return &Bot{
		token,
	}
}

func (b *Bot) Init() {
	client, err := discordgo.New("Bot " + b.token)
	if err != nil {
		return
	}
	client.AddHandler(ready)
	client.AddHandler(createMessage)
	client.Identify.Intents = discordgo.IntentsGuildMessages

	client.Open()
	fmt.Println("started")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	client.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, "lol")
	fmt.Println("ready")
}

func createMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID ||
		!strings.HasPrefix(m.Content, prefix) {
		return
	}

	if strings.HasPrefix(m.Content, prefix+"t") {
		s.ChannelMessageSend(m.ChannelID, "Started Tic Tac Toe")
		player1 := m.Author.ID
		player2 := strings.Trim(m.Content, "<@")
		player2 = strings.Trim(player2, ">")

		t := NewTicTacToe(player1, player2, s)
		games = append(games, t)
		go t.Start(s, m)
	}
}
