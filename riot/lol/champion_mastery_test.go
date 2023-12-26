package lol

import (
	"errors"
	"testing"

	"github.com/JoaoDanielRufino/goriot/internal/request/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetChampionMasteriesByPuuid(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      []ChampionMastery
		err       error
	}{
		{
			name:      "should successfully get champion masteries by puuid",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetChampionMasteriesByPuuid("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetChampionMasteryByPuuidAndChampionId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      ChampionMastery
		err       error
	}{
		{
			name:      "should successfully get champion mastery by puuid and champion id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      ChampionMastery{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      ChampionMastery{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetChampionMasteryByPuuidAndChampionId("ID", "ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTopChampionMasteriesByPuuid(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      []ChampionMastery
		err       error
	}{
		{
			name:      "should successfully get top champion masteries by puuid",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetTopChampionMasteriesByPuuid("ID", WithCount(5))

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetChampionMasteriesBySummonerId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      []ChampionMastery
		err       error
	}{
		{
			name:      "should successfully get champion masteries by summoner id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetChampionMasteriesBySummonerId("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetChampionMasteryBySummonerIdAndChampionId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      ChampionMastery
		err       error
	}{
		{
			name:      "should successfully get champion mastery by summoner id and champion id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      ChampionMastery{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      ChampionMastery{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetChampionMasteryBySummonerIdAndChampionId("ID", "ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTopChampionMasteriesBySummonerId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      []ChampionMastery
		err       error
	}{
		{
			name:      "should successfully get top champion masteries by summoner id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []ChampionMastery{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetTopChampionMasteriesBySummonerId("ID", WithCount(5))

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetChampionMasteryScoreByPuuid(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      int
		err       error
	}{
		{
			name:      "should successfully get champion mastery by puuid",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      100,
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      0,
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetChampionMasteryScoreByPuuid("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetChampionMasteryScoreBySummonerId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      int
		err       error
	}{
		{
			name:      "should successfully get champion mastery by summoner id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      100,
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      0,
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetChampionMasteryScoreBySummonerId("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
