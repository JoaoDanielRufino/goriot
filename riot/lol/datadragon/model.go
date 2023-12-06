package datadragon

type ChampionInfo struct {
	Attack     int
	Defense    int
	Magic      int
	Difficulty int
}

type ChampionImage struct {
	Full   string
	Sprite string
	Group  string
	X      int
	Y      int
	W      int
	H      int
}

type ChampionStats struct {
	HealthPoints                    float64 `json:"hp"`
	HealthPointsPerLevel            float64 `json:"hpperlevel"`
	ManaPoints                      float64 `json:"mp"`
	ManaPointsPerLevel              float64 `json:"mpperlevel"`
	MovementSpeed                   float64 `json:"movespeed"`
	Armor                           float64 `json:"armor"`
	ArmorPerLevel                   float64 `json:"armorperlevel"`
	SpellBlock                      float64 `json:"spellblock"`
	SpellBlockPerLevel              float64 `json:"spellblockperlevel"`
	AttackRange                     float64 `json:"attackrange"`
	HealthPointRegeneration         float64 `json:"hpregen"`
	HealthPointRegenerationPerLevel float64 `json:"hpregenperlevel"`
	ManaPointRegeneration           float64 `json:"mpregen"`
	ManaPointRegenerationPerLevel   float64 `json:"mpregenperlevel"`
	CriticalStrikeChance            float64 `json:"crit"`
	CriticalStrikeChancePerLevel    float64 `json:"critperlevel"`
	AttackDamage                    float64 `json:"attackdamage"`
	AttackDamagePerLevel            float64 `json:"attackdamageperlevel"`
	AttackSpeedOffset               float64 `json:"attackspeedoffset"`
	AttackSpeedPerLevel             float64 `json:"attackspeedperlevel"`
}

type Champion struct {
	Version string
	Id      string
	Key     string
	Name    string
	Title   string
	Blurb   string
	Info    ChampionInfo
	Image   ChampionImage
	Tags    []string
	Partype string
	Stats   ChampionStats
}

type DataDragonResponse[T any] struct {
	Type    string
	Format  string
	Version string
	Data    map[string]T
}
