package lol

const (
	liveClientDataBaseURL = "https://127.0.0.1:2999/liveclientdata"

	lolBaseEndpoint = "/lol"

	leagueBaseEndpoint               = lolBaseEndpoint + "/league/v4"
	leagueChallengerByQueueEndpoint  = leagueBaseEndpoint + "/challengerleagues/by-queue/%s"
	leagueGrandMasterByQueueEndpoint = leagueBaseEndpoint + "/grandmasterleagues/by-queue/%s"
	leagueMasterByQueueEndpoint      = leagueBaseEndpoint + "/masterleagues/by-queue/%s"
	leagueBySummonerIdEndpoint       = leagueBaseEndpoint + "/entries/by-summoner/%s"
	leagueEntriesEndpoint            = leagueBaseEndpoint + "/entries/%s/%s/%s"
	leagueByIdEndpoint               = leagueBaseEndpoint + "/leagues/%s"

	summonerBaseEndpoint                  = lolBaseEndpoint + "/summoner/v4/summoners"
	summonerByAccountEndpoint             = summonerBaseEndpoint + "/by-account/%s"
	summonerByNameEndpoint                = summonerBaseEndpoint + "/by-name/%s"
	summonerByPuuidEndpoint               = summonerBaseEndpoint + "/by-puuid/%s"
	summonerByEncryptedSummonerIdEndpoint = summonerBaseEndpoint + "/%s"
)
