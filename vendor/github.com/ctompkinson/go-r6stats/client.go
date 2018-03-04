package r6

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/pkg/errors"
)

func NewClient(httpClient http.Client) (client) {
	return client{httpClient}
}

func (c *client) GetPlayer (playerName string, platform string) (Player, error){
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
	return playerRes.Player, nil
}

func (c *client) GetPlayerUrl (player string, platform string) string {
	return fmt.Sprintf(Endpoint + "/players/%s?platform=%s", player, platform)
}
