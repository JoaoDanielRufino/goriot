package live

const (
	BaseURL = "https://127.0.0.1:2999/liveclientdata"

	allGameDataEndpoint           = "/allgamedata"
	activePlayerEndpoint          = "/activeplayer"
	activePlayerNameEndpoint      = "/activeplayername"
	activePlayerAbilitiesEndpoint = "/activeplayerabilities"
	activePlayerRunesEndpoint     = "/activeplayerrunes"
	playerListEndpoint            = "/playerlist"
	playerScoresEndpoint          = "/playerscores?summonerName=%s"
	playerSummonerSpellsEndpoint  = "/playersummonerspells?summonerName=%s"
	playerMainRunesEndpoint       = "/playermainrunes?summonerName=%s"
	playerItemsEndpoint           = "/playeritems?summonerName=%s"
	eventDataEndpoint             = "/eventdata"
	gameStatsEndpoint             = "/gamestats"
)
