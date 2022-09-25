package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Command-line flags.
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()
}

func main() {
	ds, err := discordgo.New(fmt.Sprintf("Bot %s", Token))
	if err != nil {
		fmt.Printf("ERROR: Could not create Discord session. %s", err)
		return
	}

	ds.Identify.Intents = discordgo.IntentsGuildMessages

	err = ds.Open()
	if err != nil {
		fmt.Printf("ERROR: Could not establish websocket connection. %s", err)
		return
	}

	fmt.Print("Hello, world!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// Block until term signal received, then gracefully shut down.
	<-c
	ds.Close()
}
