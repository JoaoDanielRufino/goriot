package lol

import (
	"fmt"
	"net/url"
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
