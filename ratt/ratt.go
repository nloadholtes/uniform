package ratt

import (
	"github.com/fabioxgn/go-bot"
	"math/rand"
	"regexp"
)

var (
	re          = regexp.MustCompile("(?i)\\b(ratt)\\b")
	ratt_lyrics = []string{
		"What goes around comes around",
		"Dance dance dance",
		"Round and round",
		"You should know, by now",
	}
)

func ratt(command *bot.PassiveCmd) (msg string, err error) {
	if re.MatchString(command.Raw) {
		return ratt_lyrics[rand.Intn(len(ratt_lyrics))], nil
	}
	return "", nil
}

func init() {
	bot.RegisterCommand("ratt", "Spouts out lyrics from the band Ratt", "whatever", ratt)
}
