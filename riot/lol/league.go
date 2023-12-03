package lol

import "fmt"

func (l *LoL) GetLeagueBySummonerId(summonerId string) ([]LeagueEntry, error) {
	leagueEntry := []LeagueEntry{}
	return leagueEntry, l.get(fmt.Sprintf(leagueBySummonerId, summonerId), &leagueEntry)
}
