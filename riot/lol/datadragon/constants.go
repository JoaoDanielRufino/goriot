package datadragon

import "github.com/JoaoDanielRufino/goriot/riot"

const (
	BaseURL = "https://ddragon.leagueoflegends.com"

	realmEndpoint = "/realms/%s.json"
)

var (
	regionToRealm = map[string]string{
		riot.RegionBrazil:            "br",
		riot.RegionEuropeWest:        "euw",
		riot.RegionEuropeNorthEast:   "eun",
		riot.RegionJapan:             "jp",
		riot.RegionKorea:             "kr",
		riot.RegionLatinAmericaNorth: "lan",
		riot.RegionLatinAmericaSouth: "las",
		riot.RegionNorthAmerica:      "na",
		riot.RegionOceania:           "oce",
		riot.RegionPBE:               "pbe",
		riot.RegionPhilippines:       "ph",
		riot.RegionRussia:            "ru",
		riot.RegionSingapore:         "sg",
		riot.RegionThailand:          "th",
		riot.RegionTurkey:            "tr",
		riot.RegionTaiwan:            "tw",
		riot.RegionVietnam:           "vn",
	}
)
