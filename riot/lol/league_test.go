package lol

import (
	"errors"
	"testing"

	"github.com/JoaoDanielRufino/goriot/internal/request/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetChallengerLeaguesByQueue(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      *LeagueList
		err       error
	}{
		{
			name:      "should successfully get challenger leagues by queue",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetChallengerLeaguesByQueue("QUEUE")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetGrandMasterLeaguesByQueue(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      *LeagueList
		err       error
	}{
		{
			name:      "should successfully get grand master leagues by queue",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetGrandMasterLeaguesByQueue("QUEUE")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetMasterLeaguesByQueue(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      *LeagueList
		err       error
	}{
		{
			name:      "should successfully get master leagues by queue",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetMasterLeaguesByQueue("QUEUE")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLeagueBySummonerId(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      []LeagueEntry
		err       error
	}{
		{
			name:      "should successfully get league by summoner id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []LeagueEntry{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []LeagueEntry{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetLeagueBySummonerId("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLeagueEntries(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      []LeagueEntry
		err       error
	}{
		{
			name:      "should successfully get league entries",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []LeagueEntry{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      []LeagueEntry{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetLeagueEntries("QUEUE", "TIER", "DIVISION")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLeagueById(t *testing.T) {
	tests := []struct {
		name      string
		lolClient *LoL
		want      *LeagueList
		err       error
	}{
		{
			name:      "should successfully get league by id",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       nil,
		},
		{
			name:      "should throw not found error",
			lolClient: &LoL{apiKey: "API_KEY"},
			want:      &LeagueList{},
			err:       errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.lolClient.httpClient = mock.JsonResponseMock(tt.want, tt.err)

			got, err := tt.lolClient.GetLeagueById("ID")

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
