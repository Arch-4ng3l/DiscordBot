package main

import "flag"

func main() {

	BotToken := flag.String("token", "", "Token for Bot")
	flag.Parse()
	bot := NewBot(*BotToken)
	bot.Init()

}
