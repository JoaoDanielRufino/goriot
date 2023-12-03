package lol

import (
	"fmt"
)

type ParamsOptions func(endpoint *string)

func WithCount(count uint) ParamsOptions {
	return func(endpoint *string) {
		*endpoint = fmt.Sprintf("%s?count=%d", *endpoint, count)
	}
}

func (l *LoL) GetChampionMasteryByPuuid(puuid string) ([]ChampionMastery, error) {
	championMastery := []ChampionMastery{}
	return championMastery, l.get(fmt.Sprintf(championMasteryByPuuidEndpoint, puuid), &championMastery)
}

func (l *LoL) GetChampionMasteryByPuuidAndChampionId(puuid, championId string) (ChampionMastery, error) {
	championMastery := ChampionMastery{}
	return championMastery, l.get(fmt.Sprintf(championMasteryByPuuidAndChampionIdEndpoint, puuid, championId), &championMastery)
}

func (l *LoL) GetTopChampionMasteryByPuuid(puuid string, options ...ParamsOptions) ([]ChampionMastery, error) {
	championMastery := []ChampionMastery{}

	endpoint := fmt.Sprintf(championMasteryByPuuidTopEndpoint, puuid)

	for _, opt := range options {
		opt(&endpoint)
	}

	return championMastery, l.get(endpoint, &championMastery)
}

func (l *LoL) GetChampionMasteryBySummonerId(summonerId string) ([]ChampionMastery, error) {
	championMastery := []ChampionMastery{}
	return championMastery, l.get(fmt.Sprintf(championMasteryBySummonerIdEndpoint, summonerId), &championMastery)
}

func (l *LoL) GetChampionMasteryBySummonerIdAndChampionId(summonerId, championId string) (ChampionMastery, error) {
	championMastery := ChampionMastery{}
	return championMastery, l.get(fmt.Sprintf(championMasteryBySummonerIdAndChampionIdEndpoint, summonerId, championId), &championMastery)
}

func (l *LoL) GetTopChampionMasteryBySummonerId(summonerId string, options ...ParamsOptions) ([]ChampionMastery, error) {
	championMastery := []ChampionMastery{}

	endpoint := fmt.Sprintf(championMasteryBySummonerIdTopEndpoint, summonerId)

	for _, opt := range options {
		opt(&endpoint)
	}

	return championMastery, l.get(endpoint, &championMastery)
}

// Player's total champion mastery score, which is the sum of every champion mastery level
func (l *LoL) GetChampionMasteryScoreByPuuid(puuid string) (int, error) {
	var score int
	return score, l.get(fmt.Sprintf(championMasteryScoresByPuuidEndpoint, puuid), &score)
}

// Player's total champion mastery score, which is the sum of every champion mastery level
func (l *LoL) GetChampionMasteryScoreBySummonerId(summonerId string) (int, error) {
	var score int
	return score, l.get(fmt.Sprintf(championMasteryScoresBySummonerIdEndpoint, summonerId), &score)
}
