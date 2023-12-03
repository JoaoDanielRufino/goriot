package lol

const (
	lolBaseEndpoint = "/lol"

	leagueBaseEndpoint = lolBaseEndpoint + "/league/v4"
	leagueBySummonerId = leagueBaseEndpoint + "/entries/by-summoner/%s"

	summonerBaseEndpoint                  = lolBaseEndpoint + "/summoner/v4/summoners"
	summonerByAccountEndpoint             = summonerBaseEndpoint + "/by-account/%s"
	summonerByNameEndpoint                = summonerBaseEndpoint + "/by-name/%s"
	summonerByPuuidEndpoint               = summonerBaseEndpoint + "/by-puuid/%s"
	summonerByEncryptedSummonerIdEndpoint = summonerBaseEndpoint + "/%s"
)
