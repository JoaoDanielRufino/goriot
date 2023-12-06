package datadragon

import (
	"encoding/json"
	"fmt"

	"github.com/JoaoDanielRufino/goriot/internal/request"
)

type DataDragon struct {
	version    string
	language   string
	cdn        string
	httpClient *request.HttpClient
}

func NewDataDragon(region string, httpClient *request.HttpClient) (*DataDragon, error) {
	if _, ok := regionToRealm[region]; !ok {
		return nil, fmt.Errorf("failed to find a realm for the region: %s", region)
	}

	datadragon := &DataDragon{
		httpClient: httpClient,
	}

	if err := datadragon.init(regionToRealm[region]); err != nil {
		return nil, err
	}

	return datadragon, nil
}

func (d *DataDragon) init(realm string) error {
	var body struct {
		V   string
		L   string
		Cnd string
	}

	if err := d.get(fmt.Sprintf(realmEndpoint, realm), &body); err != nil {
		return err
	}

	d.version = body.V
	d.language = body.L
	d.cdn = body.Cnd

	return nil
}

func (d *DataDragon) get(endpoint string, target interface{}) error {
	resp, err := d.httpClient.Do(request.GET, endpoint, nil)
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