package lol

import (
	"fmt"
	"os"

	"github.com/JoaoDanielRufino/goriot/internal"
	"github.com/JoaoDanielRufino/goriot/riot"
	"github.com/JoaoDanielRufino/goriot/riot/lol/live"
)

type LoL struct {
	LiveClientData *live.LiveClientData
	apiKey         string
	region         string
	httpClient     *internal.HttpClient
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

		l.LiveClientData = live.NewLiveClientData(internal.NewSecureHttpClient(riot.LiveClientDataBaseURL, certificate))
		return nil
	}
}

func NewClient(options ...Option) (*LoL, error) {
	lolClient := &LoL{
		LiveClientData: live.NewLiveClientData(internal.NewInsecureHttpClient(riot.LiveClientDataBaseURL)),
		region:         riot.RegionBrasil,
	}

	for _, opt := range options {
		err := opt(lolClient)
		if err != nil {
			return nil, err
		}
	}

	lolClient.httpClient = internal.NewDefaultHttpClient(fmt.Sprintf(riot.ApiBaseURL, lolClient.region))

	return lolClient, nil
}
