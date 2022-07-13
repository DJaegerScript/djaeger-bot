package main

import (
	"djaeger-bot/database"
	"djaeger-bot/handlers"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	err := godotenv.Load()
	if err != nil {
		return
	}
	DB := database.Init()
	handler := handlers.New(DB)

	dg, sessionError := discordgo.New(`Bot ` + os.Getenv("DISCORD_TOKEN"))
	if sessionError != nil {
		panic(sessionError)
	}

	dg.AddHandler(handler.Todo)

	dg.Identify.Intents = discordgo.IntentGuildMessages

	connectionError := dg.Open()

	defer func(dg *discordgo.Session) {
		err := dg.Close()
		if err != nil {

		}
	}(dg)

	if connectionError != nil {
		panic(connectionError)
	}

	fmt.Println("djaeger bot is listening to the command")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
