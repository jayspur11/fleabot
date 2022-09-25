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
		fmt.Println("ERROR: Could not create Discord session.", err)
		return
	}

	ds.Identify.Intents = discordgo.IntentsGuildMessages
	ds.AddHandler(messageCreateEvent)

	err = ds.Open()
	if err != nil {
		fmt.Println("ERROR: Could not establish websocket connection.", err)
		return
	}

	fmt.Println("Hello, world!")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// Block until term signal received, then gracefully shut down.
	<-c
	ds.Close()
}

// Event Handlers

func messageCreateEvent(ds *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("New message from %s.\n", m.Author.Username)
}
