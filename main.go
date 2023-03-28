package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvent(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println(event.Timestamp)
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5023612414387-5047503036736-YqItzBByPBLGwcSxlrivUt7j")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A050PNU2T6E-5023781968594-218084214defb6285d8be54362421eb09aecf0a6c235481127cdfdf02fbdc2dd")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvent(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "Age Calculator",
		Examples:    []string{"my yob is 2001"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				log.Fatal(err)
			}
			age := 2022 - yob
			response.Reply(fmt.Sprintf("Your age is %d", age))
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
