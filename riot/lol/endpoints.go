package lol

const (
	lolBaseEndpoint = "/lol"

	summonerBaseEndpoint                  = lolBaseEndpoint + "/summoner/v4/summoners"
	summonerByAccountEndpoint             = summonerBaseEndpoint + "/by-account/%s"
	summonerByNameEndpoint                = summonerBaseEndpoint + "/by-name/%s"
	summonerByPuuidEndpoint               = summonerBaseEndpoint + "/by-puuid/%s"
	summonerByEncryptedSummonerIdEndpoint = summonerBaseEndpoint + "/%s"
)
