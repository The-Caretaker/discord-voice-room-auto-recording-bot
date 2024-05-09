package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	sess, err := discordgo.New("Bot MTIzODAyOTYwODA2NzI3MjcwNQ.GYq0L9.2n4JWR4u4Wdtpwx6uo5Ebd4o2sPxU2Ot1OCnoU")
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) /*Hello World debug handler*/ {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "нимагу работать" {
			s.ChannelMessageSend(m.ChannelID, "балею)")
		}
	})

	sess.AddHandler(func(s *discordgo.Session, event *discordgo.ChannelCreate) /*Join Voice Channel*/ {
		_, err := sess.ChannelVoiceJoin(event.GuildID, event.ID, false, false)
		if err != nil {
			log.Fatal(err)
		}
	})

	sess.AddHandler(func(s *discordgo.Session, event *discordgo.ChannelDelete) {
		sess.ThreadLeave(event.ID)
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
