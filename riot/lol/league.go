package lol

import "fmt"

func (l *LoL) GetChallengerLeaguesByQueue(queue string) (*LeagueList, error) {
	leagueList := &LeagueList{}
	return leagueList, l.get(fmt.Sprintf(leagueChallengerByQueueEndpoint, queue), leagueList)
}

func (l *LoL) GetGrandMasterLeaguesByQueue(queue string) (*LeagueList, error) {
	leagueList := &LeagueList{}
	return leagueList, l.get(fmt.Sprintf(leagueGrandMasterByQueueEndpoint, queue), leagueList)
}

func (l *LoL) GetMasterLeaguesByQueue(queue string) (*LeagueList, error) {
	leagueList := &LeagueList{}
	return leagueList, l.get(fmt.Sprintf(leagueMasterByQueueEndpoint, queue), leagueList)
}

func (l *LoL) GetLeagueBySummonerId(summonerId string) ([]LeagueEntry, error) {
	leagueEntry := []LeagueEntry{}
	return leagueEntry, l.get(fmt.Sprintf(leagueBySummonerIdEndpoint, summonerId), &leagueEntry)
}

func (l *LoL) GetLeagueEntries(queue, tier, division string) ([]LeagueEntry, error) {
	leagueEntry := []LeagueEntry{}
	return leagueEntry, l.get(fmt.Sprintf(leagueEntriesEndpoint, queue, tier, division), &leagueEntry)
}

func (l *LoL) GetLeagueById(leagueId string) (*LeagueList, error) {
	leagueList := &LeagueList{}
	return leagueList, l.get(fmt.Sprintf(leagueByIdEndpoint, leagueId), leagueList)
}
