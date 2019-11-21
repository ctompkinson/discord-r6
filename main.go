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

func main() {
	Token = os.Getenv("DISCORD_TOKEN")
	if Token == "" {
		log.Println("Token not found please set env variable DISCORD_TOKEN")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(getMessageHandler)

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
func getMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// If we didn't receive an /r6 command then we don't need to continue
	if !strings.Contains(m.Message.Content, "/r6") {
		return
	}

	// split the message up for processing
	parts := strings.Split(m.Message.Content, " ")

	if len(parts) == 1 {
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, newMsg("No player specified"))
		return
	}

	// Send the initial loading message
	fmt.Println("Sending stats")
	loadingMsg := discordgo.MessageEmbed{
		Title: "Loading stats for " + parts[1],
	}
	msg, err := s.ChannelMessageSendEmbed(m.ChannelID, &loadingMsg)
	if err != nil {
		log.Println(err)
		return
	}

	// Create the r6stats client and fetch the player
	r6 := r6.NewClient(http.Client{})
	player, err := r6.GetPlayer(parts[1], "uplay", true)
	if err != nil {
		_, err = s.ChannelMessageEditEmbed(m.ChannelID, m.Message.ID, &discordgo.MessageEmbed{
			Title: "Failed to load player",
		})
		return
	}

	if len(parts) == 2 {
		// /r6 ruffbabe
		// Print the player stats
		sendStatsMessage(s, m, player, msg)
	} else if len(parts) == 3 {
		// /r6 ruffbabe fuze
		// Print the player operator stats
		sendOperatorMessage(s, m, player, strings.TrimSpace(parts[2]), msg)
	}

	return
}

func sendStatsMessage(s *discordgo.Session, m *discordgo.MessageCreate, player r6.Player, msg *discordgo.Message) {
	favOff, favDef := favOps(player)

	favOffKD := operatorKD(&favOff)
	favOffWL := operatorWL(&favOff)

	favDefKD := operatorKD(&favDef)
	favDefWL := operatorWL(&favDef)

	statsMsg := discordgo.MessageEmbed{
		Title: fmt.Sprintf("Siege Stats for %s", player.Username),
		Color: 10,
		Fields: []*discordgo.MessageEmbedField{
			// Ranked
			{
				Name:   "Ranked PlayTime",
				Value:  strconv.FormatFloat(player.Stats.Ranked.Playtime/60/60, 'f', 1, 64),
				Inline: true,
			},
			{
				Name:   "Kill/Death",
				Value:  strconv.FormatFloat(player.Stats.Ranked.KillDeathRatio, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Win/Loss",
				Value:  strconv.FormatFloat(player.Stats.Ranked.WinLossRatio, 'f', 3, 64),
				Inline: true,
			},

			// Casual
			{
				Name:   "Casual PlayTime",
				Value:  strconv.FormatFloat(player.Stats.Casual.Playtime/60/60, 'f', 1, 64),
				Inline: true,
			},
			{
				Name:   "Kill/Death",
				Value:  strconv.FormatFloat(player.Stats.Casual.KillDeathRatio, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Win/Loss",
				Value:  strconv.FormatFloat(player.Stats.Casual.WinLossRatio, 'f', 3, 64),
				Inline: true,
			},

			// Offensive Operator
			{
				Name:   "Favorite Atk",
				Value:  favOff.Name,
				Inline: true,
			},
			{
				Name:   "Kill/Death",
				Value:  strconv.FormatFloat(favOffKD, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Win/Loss",
				Value:  strconv.FormatFloat(favOffWL, 'f', 3, 64),
				Inline: true,
			},

			// Defensive Operator
			{
				Name:   "Favorite Def",
				Value:  favDef.Name,
				Inline: true,
			},
			{
				Name:   "Kill/Death",
				Value:  strconv.FormatFloat(favDefKD, 'f', 3, 64),
				Inline: true,
			},
			{
				Name:   "Win/Loss",
				Value:  strconv.FormatFloat(favDefWL, 'f', 3, 64),
				Inline: true,
			},
		},
	}

	s.ChannelMessageEditEmbed(m.ChannelID, msg.ID, &statsMsg)
}

func sendOperatorMessage(s *discordgo.Session, m *discordgo.MessageCreate, player r6.Player, operatorName string, msg *discordgo.Message) {
	fmt.Println("getting stats for " + operatorName)

	opName := strings.Title(operatorName)

	if _, ok := player.Operators[opName]; !ok {
		_, err := s.ChannelMessageEditEmbed(msg.ChannelID, msg.ID, &discordgo.MessageEmbed{
			Title: "Invalid operator name",
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	op := player.Operators[opName]

	// Generate the fields
	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "PlayTime",
			Value:  strconv.FormatFloat(float64(op.Playtime)/60/60, 'f', 1, 64),
			Inline: true,
		},
		{
			Name:
			"Kill/Death",
			Value:  strconv.FormatFloat(operatorKD(&op), 'f', 3, 64),
			Inline: true,
		},
		{
			Name:   "Win/Loss",
			Value:  strconv.FormatFloat(operatorWL(&op), 'f', 3, 64),
			Inline: true,
		},
	}

	fmt.Println(op.Specials)
	for k, v := range op.Specials {
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   k,
			Value:  v,
			Inline: true,
		})
	}

	opMsg := discordgo.MessageEmbed{
		Title: fmt.Sprintf("Siege Stats for %s: %s", player.Username, op.Name),
		Color: 10,
		Fields: fields,
	}

	s.ChannelMessageEditEmbed(m.ChannelID, msg.ID, &opMsg)
}

func operatorKD(op *r6.Operator) float64 {
	return float64(op.Kills) / float64(op.Deaths)
}

func operatorWL(op *r6.Operator) float64 {
	return float64(op.Wins) / float64(op.Losses)
}

func favOps(player r6.Player) (favOff r6.Operator, favDef r6.Operator) {
	// Some intial operators to compare against
	favOff = player.Operators["ash"]
	favDef = player.Operators["frost"]

	for _, op := range player.Operators {
		if op.Role == "atk" {
			if op.Playtime > favOff.Playtime {
				favOff = op
			}
		}

		if op.Role == "def" {
			if op.Playtime > favDef.Playtime {
				favDef = op
			}
		}
	}

	return
}

func newMsg(description string) *discordgo.MessageEmbed {
	msg := discordgo.MessageEmbed{
		Title:       "Siege Stats",
		Color:       2,
		Description: description,
	}

	return &msg
}
