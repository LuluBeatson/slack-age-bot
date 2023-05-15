package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed to load .env file:", err)
		return
	}
}

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{
		Description: "Year of birth calculator",
		Examples:    []string{"My year of birth is 1980"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			y, err := strconv.Atoi(year)
			if err != nil {
				response.Reply("Invalid year")
				return
			}
			// get current year
			current_year := time.Now().Year()
			response.Reply(fmt.Sprintf("Your age is %d", current_year-y))
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(commandEvents <-chan *slacker.CommandEvent) {
	for event := range commandEvents {
		fmt.Println("Command Event. Timestamp:", event.Timestamp, "Command:", event.Command, "Parameters:", event.Parameters)
	}
}
