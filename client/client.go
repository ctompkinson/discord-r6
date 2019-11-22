package client

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func NewClient(httpClient http.Client) client {
	return client{httpClient}
}

func (c *client) GetPlayer(playerName string, platform string, operators bool) (Player, error) {
	playerUrl := c.GetPlayerUrl(playerName, platform)
	res, err := c.httpClient.Get(playerUrl)

	if err != nil {
		return Player{}, errors.Wrap(err, "unable to retrieve player")
	}

	if res.StatusCode != http.StatusOK {
		return Player{}, errors.New(
			fmt.Sprintf("got status %s when retrieving player", strconv.Itoa(res.StatusCode)),
		)
	}

	decoder := json.NewDecoder(res.Body)
	var playerRes PlayerResponse
	resErr := decoder.Decode(&playerRes)
	if resErr != nil {
		fmt.Println(resErr)
		return Player{}, errors.Wrap(err, "Unable to retrieve player")
	}

	res.Body.Close()

	// If we're getting the operator information to an additional get for the operator information
	if operators {
		c.getOperators(playerName, platform, &playerRes.Player)
	}

	return playerRes.Player, nil
}

// Takes an existing player and populates it with operator information
func (c *client) getOperators(playerName string, platform string, player *Player) error {
	operatorUrl := c.GetPlayerOperatorsUrl(playerName, platform)
	res, err := c.httpClient.Get(operatorUrl)

	if err != nil {
		return errors.Wrap(err, "unable to retrieve player")
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("got status %s when retrieving operators", strconv.Itoa(res.StatusCode)))
	}

	var operatorResponse OperatorResponse
	decoder := json.NewDecoder(res.Body)
	resErr := decoder.Decode(&operatorResponse)
	if resErr != nil {
		fmt.Println(resErr)
		return errors.Wrap(err, "Unable to retrieve operators")
	}
	res.Body.Close()

	for _, op := range operatorResponse.OperatorRecords {
		if player.Operators == nil {
			player.Operators = map[string]Operator{}
		}

		player.Operators[op.Operator.Name] = Operator{
			Name:     op.Operator.Name,
			Role:     op.Operator.Role,
			Wins:     op.Stats.Wins,
			Losses:   op.Stats.Losses,
			Kills:    op.Stats.Kills,
			Deaths:   op.Stats.Deaths,
			Playtime: op.Stats.Playtime,
			Specials: op.Stats.Specials,
			Images:   op.Operator.Images,
		}
	}

	return nil
}

func (c *client) GetPlayerUrl(player string, platform string) string {
	return fmt.Sprintf(Endpoint+"/players/%s?platform=%s", player, platform)
}

func (c *client) GetPlayerOperatorsUrl(player string, platform string) string {
	return fmt.Sprintf(Endpoint+"/players/%s/operators?platform=%s", player, platform)
}
