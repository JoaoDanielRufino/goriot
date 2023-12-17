package lol

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/JoaoDanielRufino/goriot/internal"
	"github.com/JoaoDanielRufino/goriot/internal/request"
	"github.com/JoaoDanielRufino/goriot/riot"
	"github.com/JoaoDanielRufino/goriot/riot/lol/datadragon"
	"github.com/JoaoDanielRufino/goriot/riot/lol/live"
)

type LoL struct {
	LiveClientData *live.LiveClientData
	DataDragon     *datadragon.DataDragon
	httpClient     request.Requester
	apiKey         string
	region         string
}

type Option func(*LoL) error

func WithApiKey(key string) Option {
	return func(l *LoL) error {
		l.apiKey = key
		return nil
	}
}

func WithRegion(region string) Option {
	return func(l *LoL) error {
		l.region = region
		return nil
	}
}

func WithCertificate(certificatePath string) Option {
	return func(l *LoL) error {
		certificate, err := os.ReadFile(certificatePath)
		if err != nil {
			return err
		}

		l.LiveClientData = live.NewLiveClientData(request.NewSecureHttpClient(live.BaseURL, certificate))
		return nil
	}
}

func NewClient(options ...Option) (*LoL, error) {
	lolClient := &LoL{
		LiveClientData: live.NewLiveClientData(request.NewInsecureHttpClient(live.BaseURL)),
		region:         riot.RegionBrazil,
	}

	for _, opt := range options {
		if err := opt(lolClient); err != nil {
			return nil, err
		}
	}

	lolClient.httpClient = request.NewDefaultHttpClient(fmt.Sprintf(internal.ApiBaseURL, lolClient.region))

	dataDragonClient, err := datadragon.NewDataDragon(lolClient.region, request.NewDefaultHttpClient(datadragon.BaseURL))
	if err != nil {
		return nil, err
	}

	lolClient.DataDragon = dataDragonClient

	return lolClient, nil
}

func (l *LoL) get(endpoint string, target interface{}) error {
	resp, err := l.httpClient.Do(request.GET, endpoint, nil, request.WithApiKey(l.apiKey))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}
