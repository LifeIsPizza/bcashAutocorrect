package main

import (
	"log"
	"os"
	"strings"

	"github.com/yanzay/tbot"
)

func main() {
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	bot.HandleFunc("/correct {text}", BcashHandler)
	err = bot.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func BcashHandler(m *tbot.Message) {
	//reply := fmt.Sprintf("```\n%s\n```", bcashCorrect(m.Vars["text"]))
	//m.Reply(reply, tbot.WithMarkdown)
	m.Reply(bcashCorrect(m.Vars["text"]), tbot.WithMarkdown)

}

func bcashCorrect(text string) string {
	//correctedText := strings.Replace(text, "Bitcoin Cash", "Bcash", -1)
	r := strings.NewReplacer("Bitcoin Cash", "bcash",
		"bitcoin cash", "bcash",
		"Bitcoin cash", "bcash",
		"Bitcoin (BCH)", "bcash",
		"bitcoincash", "bcash",
		"BitcoinCash", "bcash")
	//textLine := fmt.Sprintf("%s", text)
	correctedText := r.Replace(text)
	return correctedText
}
