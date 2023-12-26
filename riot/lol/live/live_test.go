package live

import (
	"errors"
	"testing"

	"github.com/JoaoDanielRufino/goriot/internal/request/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewLiveClientData(t *testing.T) {
	expected := &LiveClientData{
		httpClient: mock.RequesterMock{},
	}

	got := NewLiveClientData(mock.RequesterMock{})

	assert.Equal(t, expected, got)
}

func TestAllGameData(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *AllData
		err        error
	}{
		{
			name:       "should successfully get all game data",
			liveClient: &LiveClientData{},
			want:       &AllData{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &AllData{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.AllGameData()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestActivePlayer(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *ActivePlayer
		err        error
	}{
		{
			name:       "should successfully get active player",
			liveClient: &LiveClientData{},
			want:       &ActivePlayer{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &ActivePlayer{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.ActivePlayer()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestActivePlayerName(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       string
		err        error
	}{
		{
			name:       "should successfully get active player name",
			liveClient: &LiveClientData{},
			want:       "Hide on bush",
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       "",
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.ActivePlayerName()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestActivePlayerAbilities(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *Abilities
		err        error
	}{
		{
			name:       "should successfully get active player abilities",
			liveClient: &LiveClientData{},
			want:       &Abilities{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &Abilities{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.ActivePlayerAbilities()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestActivePlayerRunes(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *FullRunes
		err        error
	}{
		{
			name:       "should successfully get active player runes",
			liveClient: &LiveClientData{},
			want:       &FullRunes{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &FullRunes{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.ActivePlayerRunes()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlayerList(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       []Player
		err        error
	}{
		{
			name:       "should successfully get player list",
			liveClient: &LiveClientData{},
			want:       []Player{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       []Player{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.PlayerList()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlayerScores(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *Scores
		err        error
	}{
		{
			name:       "should successfully get player scores",
			liveClient: &LiveClientData{},
			want:       &Scores{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &Scores{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.PlayerScores("Hide on bush")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlayerSummonerSpells(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *SummonerSpells
		err        error
	}{
		{
			name:       "should successfully get player summoner spells",
			liveClient: &LiveClientData{},
			want:       &SummonerSpells{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &SummonerSpells{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.PlayerSummonerSpells("Hide on bush")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlayerMainRunes(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *PlayerRunes
		err        error
	}{
		{
			name:       "should successfully get player main runes",
			liveClient: &LiveClientData{},
			want:       &PlayerRunes{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &PlayerRunes{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.PlayerMainRunes("Hide on bush")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlayerItems(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       []Item
		err        error
	}{
		{
			name:       "should successfully get player items",
			liveClient: &LiveClientData{},
			want:       []Item{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       []Item{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.PlayerItems("Hide on bush")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEventData(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *Events
		err        error
	}{
		{
			name:       "should successfully get event data",
			liveClient: &LiveClientData{},
			want:       &Events{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &Events{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.EventData()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGameStats(t *testing.T) {
	tests := []struct {
		name       string
		liveClient *LiveClientData
		want       *GameStats
		err        error
	}{
		{
			name:       "should successfully get game stats",
			liveClient: &LiveClientData{},
			want:       &GameStats{},
			err:        nil,
		},
		{
			name:       "should throw not found error",
			liveClient: &LiveClientData{},
			want:       &GameStats{},
			err:        errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.liveClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.liveClient.GameStats()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
