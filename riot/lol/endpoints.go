package lol

const (
	liveClientDataBaseURL = "https://127.0.0.1:2999/liveclientdata"

	lolBaseEndpoint = "/lol"

	championMasteryBaseEndpoint                      = lolBaseEndpoint + "/champion-mastery/v4"
	championMasteryByPuuidEndpoint                   = championMasteryBaseEndpoint + "/champion-masteries/by-puuid/%s"
	championMasteryByPuuidAndChampionIdEndpoint      = championMasteryBaseEndpoint + "/champion-masteries/by-puuid/%s/by-champion/%s"
	championMasteryByPuuidTopEndpoint                = championMasteryBaseEndpoint + "/champion-masteries/by-puuid/%s/top"
	championMasteryBySummonerIdEndpoint              = championMasteryBaseEndpoint + "/champion-masteries/by-summoner/%s"
	championMasteryBySummonerIdAndChampionIdEndpoint = championMasteryBaseEndpoint + "/champion-masteries/by-summoner/%s/by-champion/%s"
	championMasteryBySummonerIdTopEndpoint           = championMasteryBaseEndpoint + "/champion-masteries/by-summoner/%s/top"
	championMasteryScoresByPuuidEndpoint             = championMasteryBaseEndpoint + "/scores/by-puuid/%s"
	championMasteryScoresBySummonerIdEndpoint        = championMasteryBaseEndpoint + "/scores/by-summoner/%s"

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
