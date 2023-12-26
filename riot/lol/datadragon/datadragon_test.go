package datadragon

import (
	"errors"
	"fmt"
	"testing"

	"github.com/JoaoDanielRufino/goriot/internal/request/mock"
	"github.com/JoaoDanielRufino/goriot/riot"
	"github.com/stretchr/testify/assert"
)

type body struct {
	V string
	L string
}

func TestNewDataDragon(t *testing.T) {
	body := body{
		V: "VERSION",
		L: "LANGUAGE",
	}

	tests := []struct {
		name   string
		region string
		want   *DataDragon
		err    error
	}{
		{
			name:   "should successfully create new data dragon client",
			region: riot.RegionBrazil,
			want: &DataDragon{
				version:  body.V,
				language: body.L,
			},
			err: nil,
		},
		{
			name:   "should throw 'failed to find realm for region error' when no regionToReam match found",
			region: "REGION",
			want:   nil,
			err:    fmt.Errorf(regionToRealmException, "REGION"),
		},
		{
			name:   "should throw not found error",
			region: riot.RegionBrazil,
			want:   nil,
			err:    errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := mock.JsonResponseMock(body, tt.err)

			if tt.want != nil {
				tt.want.httpClient = httpClient
			}

			got, err := NewDataDragon(tt.region, httpClient)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetChampions(t *testing.T) {
	dataDragonResponse := DataDragonResponse[Champion]{
		Data: map[string]Champion{
			"Vladimir": Champion{},
		},
	}

	tests := []struct {
		name             string
		dataDragonClient *DataDragon
		want             []Champion
		err              error
	}{
		{
			name: "should successfully get champions",
			dataDragonClient: &DataDragon{
				version:  "VERSION",
				language: "LANGUAGE",
			},
			want: []Champion{Champion{}},
			err:  nil,
		},
		{
			name: "should throw not found error",
			dataDragonClient: &DataDragon{
				version:  "VERSION",
				language: "LANGUAGE",
			},
			want: nil,
			err:  errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dataDragonClient.httpClient = mock.JsonResponseMock(dataDragonResponse, tt.err)

			got, err := tt.dataDragonClient.GetChampions()

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
