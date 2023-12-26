package live

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JoaoDanielRufino/goriot/internal/request"
)

type LiveClientData struct {
	httpClient request.Requester
}

func NewLiveClientData(httpClient request.Requester) *LiveClientData {
	return &LiveClientData{
		httpClient: httpClient,
	}
}

func (c *LiveClientData) AllGameData() (*AllData, error) {
	allData := &AllData{}
	return allData, c.get(allGameDataEndpoint, allData)
}

func (c *LiveClientData) ActivePlayer() (*ActivePlayer, error) {
	activePlayer := &ActivePlayer{}
	return activePlayer, c.get(activePlayerEndpoint, activePlayer)
}

func (c *LiveClientData) ActivePlayerName() (string, error) {
	var name string
	return name, c.get(activePlayerNameEndpoint, &name)
}

func (c *LiveClientData) ActivePlayerAbilities() (*Abilities, error) {
	abilities := &Abilities{}
	return abilities, c.get(activePlayerAbilitiesEndpoint, abilities)
}

func (c *LiveClientData) ActivePlayerRunes() (*FullRunes, error) {
	runes := &FullRunes{}
	return runes, c.get(activePlayerRunesEndpoint, runes)
}

func (c *LiveClientData) PlayerList() ([]Player, error) {
	players := []Player{}
	return players, c.get(playerListEndpoint, &players)
}

func (c *LiveClientData) PlayerScores(name string) (*Scores, error) {
	scores := &Scores{}
	return scores, c.get(fmt.Sprintf(playerScoresEndpoint, url.QueryEscape(name)), scores)
}

func (c *LiveClientData) PlayerSummonerSpells(name string) (*SummonerSpells, error) {
	summonerSpells := &SummonerSpells{}
	return summonerSpells, c.get(fmt.Sprintf(playerSummonerSpellsEndpoint, url.QueryEscape(name)), summonerSpells)
}

func (c *LiveClientData) PlayerMainRunes(name string) (*PlayerRunes, error) {
	playerRunes := &PlayerRunes{}
	return playerRunes, c.get(fmt.Sprintf(playerMainRunesEndpoint, url.QueryEscape(name)), playerRunes)
}

func (c *LiveClientData) PlayerItems(name string) ([]Item, error) {
	items := []Item{}
	return items, c.get(fmt.Sprintf(playerItemsEndpoint, url.QueryEscape(name)), &items)
}

func (c *LiveClientData) EventData() (*Events, error) {
	events := &Events{}
	return events, c.get(eventDataEndpoint, &events)
}

func (c *LiveClientData) GameStats() (*GameStats, error) {
	gameStats := &GameStats{}
	return gameStats, c.get(gameStatsEndpoint, &gameStats)
}

func (c *LiveClientData) get(endpoint string, target interface{}) error {
	resp, err := c.httpClient.Do(request.GET, endpoint, nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}
