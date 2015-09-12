package main

import (
    "github.com/fabioxgn/go-bot"
    _ "github.com/fabioxgn/go-bot/commands/catfacts"
    _ "github.com/fabioxgn/go-bot/commands/catgif"
    _ "github.com/fabioxgn/go-bot/commands/chucknorris"
   // Import all the commands you wish to use
   "log"
   "strings"
)

func main() {
        bot.Run(&bot.Config{
            Server:   os.Getenv("IRC_SERVER"),
            Channels: strings.Split(os.Getenv("IRC_CHANNELS"), ","),
            User:     os.Getenv("IRC_USER"),
            Nick:     os.Getenv("IRC_NICK"),
            Password: os.Getenv("IRC_PASSWORD"),
            UseTLS:   true,
            Debug:    os.Getenv("DEBUG") != "",}
        )
    }
