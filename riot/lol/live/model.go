package live

type Passive struct {
	Id             string
	DisplayName    string
	RawDescription string
	RawDisplayName string
}

type Ability struct {
	Passive
	AbilityLevel uint
}

type Abilities struct {
	Passive Passive
	Q       Ability
	W       Ability
	E       Ability
	R       Ability
}

type ChampionStats struct {
	AbilityHaste                 float64
	AbilityPower                 float64
	Armor                        float64
	ArmorPenetrationFlat         float64
	ArmorPenetrationPercent      float64
	AttackDamage                 float64
	AttackRange                  float64
	AttackSpeed                  float64
	BonusArmorPenetrationPercent float64
	BonusMagicPenetrationPercent float64
	CritChance                   float64
	CritDamage                   float64
	CurrentHealth                float64
	HealShieldPower              float64
	HealthRegenRate              float64
	LifeSteal                    float64
	MagicLethality               float64
	MagicPenetrationFlat         float64
	MagicPenetrationPercent      float64
	MagicResist                  float64
	MaxHealth                    float64
	MoveSpeed                    float64
	Omnivamp                     float64
	PhysicalLethality            float64
	PhysicalVamp                 float64
	ResourceMax                  float64
	ResourceRegenRate            float64
	ResourceType                 string
	ResourceValue                float64
	SpellVamp                    float64
	Tenacity                     float64
}

type RuneDescription struct {
	Id             uint
	DisplayName    string
	RawDescription string
	RawDisplayName string
}

type StatRune struct {
	Id             uint
	RawDescription string
}

type FullRunes struct {
	GeneralRunes      []RuneDescription
	Keystone          RuneDescription
	PrimaryRuneTree   RuneDescription
	SecondaryRuneTree RuneDescription
	StatRunes         []StatRune
}

type ActivePlayer struct {
	Abilities          Abilities
	ChampionStats      *ChampionStats
	CurrentGold        float64
	FullRunes          *FullRunes
	Level              uint
	SummonerName       string
	TeamRelativeColors bool
}

type Item struct {
	CanUse         bool
	Consumable     bool
	Count          uint
	DisplayName    string
	ItemID         uint
	Price          uint
	RawDescription string
	RawDisplayName string
	Slot           uint
}

type Scores struct {
	Assists    uint
	CreepScore uint
	Deaths     uint
	Kills      uint
	WardScore  float64
}

type Spell struct {
	DisplayName    string
	RawDescription string
	RawDisplayName string
}

type SummonerSpells struct {
	SummonerSpellOne Spell
	SummonerSpellTwo Spell
}

type PlayerRunes struct {
	Keystone          RuneDescription
	PrimaryRuneTree   RuneDescription
	SecondaryRuneTree RuneDescription
}

type Player struct {
	ChampionName    string
	IsBot           bool
	IsDead          bool
	Items           []Item
	Level           uint
	Position        string
	RawChampionName string
	RawSkinName     string
	RespawnTimer    float64
	Runes           PlayerRunes
	Scores          Scores
	SkinID          uint
	SkinName        string
	SummonerName    string
	SummonerSpells  SummonerSpells
	Team            string
}

type Event struct {
	EventID   uint
	EventName string
	EventTime float64
}

type Events struct {
	Events []Event
}

type GameStats struct {
	GameMode   string
	GameTime   float64
	MapName    string
	MapNumber  uint
	MapTerrain string
}

type AllData struct {
	ActivePlayer *ActivePlayer
	AllPlayers   []*Player
	Events       Events
	GameData     GameStats
}
