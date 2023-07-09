package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TicTacToe struct {
	field         []string
	player1       string
	player2       string
	gameOver      bool
	currentPlayer uint8
}

var sym = [2]string{"X", "O"}

func NewTicTacToe(player1, player2 string, s *discordgo.Session) *TicTacToe {
	field := make([]string, 9)
	t := &TicTacToe{
		field,
		player1,
		player2,
		false,
		0,
	}
	s.AddHandler(t.Start)
	return t
}

func (t *TicTacToe) Start(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("hallo")
	if t.gameOver {

		return
	}

	if !strings.HasPrefix(m.Content, "!") {
		return
	}
	str := m.Content
	cleantStr := strings.Trim(str, "!")

	num, err := strconv.Atoi(cleantStr)

	if err != nil || num > 9 {
		return
	}

	if m.Author.ID == t.player1 && t.currentPlayer == 0 {

	} else if m.Author.ID == t.player2 && t.currentPlayer == 1 {

	} else {
		return
	}

	if t.field[num-1] == "" {
		t.field[num-1] = sym[t.currentPlayer]
	}
	t.currentPlayer = (t.currentPlayer + 1) % 2
	for i := 0; i < 6; i += 3 {
		s.ChannelMessageSend(m.ChannelID, t.field[i]+t.field[i+1]+t.field[i+2])
	}
	t.check()
}

func (t *TicTacToe) check() {
	for i := 0; i < 6; i += 3 {
		if t.field[i] == t.field[i+1] && t.field[i] == t.field[i+2] {
			t.gameOver = true
			return
		}
	}
	for i := 0; i < 3; i++ {
		if t.field[i] == t.field[i+3] && t.field[i] == t.field[i+6] {
			t.gameOver = true
			return
		}
	}
	if t.field[0] == t.field[4] && t.field[0] == t.field[8] {
		t.gameOver = true
		return
	}

	if t.field[2] == t.field[4] && t.field[2] == t.field[6] {
		t.gameOver = true
	}
}
