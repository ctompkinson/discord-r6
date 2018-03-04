package main

import (
	"log"
	"flag"
	"github.com/bwmarrin/discordgo"
	"github.com/ctompkinson/go-r6stats"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"net/http"
	"fmt"
	"strconv"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	// Create a new Discord session using the provided bot token.
	fmt.Println(Token)
	dg, err := discordgo.New("Bot " + Token)
	dg.LogLevel = 2
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(printStats)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func printStats(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.Contains(m.Message.Content, "!stats ") {
		return
	}

	parts := strings.Split(m.Message.Content, " ")
	if len(parts) == 1 {
		s.ChannelMessageSendEmbed(m.ChannelID, newMsg("No player specified"))
		return
	}

	r6 := r6.NewClient(http.Client{})
	player, err := r6.GetPlayer(parts[1], "uplay")
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, newMsg("Unable to find player"))
		return
	}

	msg := discordgo.MessageEmbed{
		Title: fmt.Sprintf("Siege Stats for %s", player.Username),
		Color: 10,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Ranked K/D",
				Value:  strconv.FormatFloat(player.Stats.Ranked.KillDeathRatio, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Ranked W/L",
				Value:  strconv.FormatFloat(player.Stats.Ranked.WinLossRatio, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Ranked PlayTime",
				Value:  strconv.FormatFloat(player.Stats.Ranked.Playtime / 60 / 60, 'f', 1, 64),
				Inline: true,
			},
			{
				Name:   "Casual K/D",
				Value:  strconv.FormatFloat(player.Stats.Casual.KillDeathRatio, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Casual W/L",
				Value:  strconv.FormatFloat(player.Stats.Casual.WinLossRatio, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Casual PlayTime",
				Value:  strconv.FormatFloat(player.Stats.Casual.Playtime / 60 / 60, 'f', 1, 64),
				Inline: true,
			},
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &msg)
}

func newMsg(description string) *discordgo.MessageEmbed {
	msg := discordgo.MessageEmbed{
		Title:       "Siege Stats",
		Color:       2,
		Description: description,
	}

	return &msg
}
