package lol

import (
	"errors"
	"testing"

	"github.com/JoaoDanielRufino/goriot/internal/request/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetSummonerByAccountId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      Summoner
		err       error
	}{
		{
			name:      "should successfully get summoner by account id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetSummonerByAccountId("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetSummonerByName(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      Summoner
		err       error
	}{
		{
			name:      "should successfully get summoner by name",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetSummonerByName("NAME")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetSummonerByPuuid(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      Summoner
		err       error
	}{
		{
			name:      "should successfully get summoner by puuid",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetSummonerByPuuid("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetSummonerBySummonerId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      Summoner
		err       error
	}{
		{
			name:      "should successfully get summoner by summoner id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      Summoner{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetSummonerBySummonerId("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
