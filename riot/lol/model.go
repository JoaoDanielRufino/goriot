package lol

type MiniSeries struct {
	Losses   uint
	Progress string
	Target   int
	Wins     uint
}

type LeagueEntry struct {
	LeagueId     string
	SummonerId   string
	QueueType    string
	Tier         string
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

type Summoner struct {
	AccountId     string
	ProfileIconId uint
	RevisionDate  uint64
	Name          string
	Id            string
	Puuid         string
	SummonerLevel uint
}
