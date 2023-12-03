package lol

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/JoaoDanielRufino/goriot/internal/request"
)

func (l *LoL) GetSummonerByAccountId(accountId string) (Summoner, error) {
	summoner := Summoner{}
	return summoner, l.get(fmt.Sprintf(summonerByAccountEndpoint, accountId), &summoner)
}

func (l *LoL) GetSummonerByName(name string) (Summoner, error) {
	summoner := Summoner{}
	return summoner, l.get(fmt.Sprintf(summonerByNameEndpoint, url.QueryEscape(name)), &summoner)
}

func (l *LoL) GetSummonerByPuuid(puuid string) (Summoner, error) {
	summoner := Summoner{}
	return summoner, l.get(fmt.Sprintf(summonerByPuuidEndpoint, puuid), &summoner)
}

func (l *LoL) GetSummonerBySummonerId(summonerId string) (Summoner, error) {
	summoner := Summoner{}
	return summoner, l.get(fmt.Sprintf(summonerByEncryptedSummonerIdEndpoint, summonerId), &summoner)
}

func (l *LoL) get(endpoint string, target interface{}) error {
	resp, err := l.httpClient.Do(request.GET, endpoint, nil, request.WithApiKey(l.apiKey))
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
