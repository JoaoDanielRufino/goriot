package lol

type MiniSeries struct {
	Losses   uint
	Progress string
	Target   int
	Wins     uint
}

type LeagueItem struct {
	SummonerName string
	SummonerId   string
	Rank         string
	LeaguePoints int
	Wins         uint
	Losses       uint
	HotStreak    bool
	Veteran      bool
	FreshBlood   bool
	Inactive     bool
	MiniSeries   *MiniSeries `json:",omitempty"`
}

type LeagueEntry struct {
	LeagueItem
	LeagueId  string
	QueueType string
	Tier      string
}

type LeagueList struct {
	LeagueId string
	Tier     string
	Name     string
	Queue    string
	Entries  []LeagueItem
}

type Summoner struct {
	AccountId     string
	ProfileIconId uint
	RevisionDate  uint64
	Name          string
	Id            string
	Puuid         string
	SummonerLevel uint
}
