# Riot Games SDK for Go

goriot is a SDK written in Go to facilitate the use of Riot Games API. Currently the SDK implements the League of Legends, Live Client Data and Data Dragon API. 

## Installation

```bash
$ go get github.com/JoaoDanielRufino/goriot
```

## Usage

To use League of Legends API, you need to provide your Riot Games API Key. You can generate a development key by accessing [developer.riotgames.com](https://developer.riotgames.com/)

The default region of the SDK is Brazil (br1).

If you only want to use Live Client Data or Data Dragon, the API Key is not required.

### Basic example

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/JoaoDanielRufino/goriot/riot"
	"github.com/JoaoDanielRufino/goriot/riot/lol"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := lol.NewClient(lol.WithApiKey("API KEY"), lol.WithRegion(riot.RegionKorea))
	handleError(err)

	res, err := client.GetSummonerByName("Hide on bush")
	handleError(err)

	pretty, err := json.MarshalIndent(res, "", "  ")
	handleError(err)

	fmt.Println(string(pretty))
}
```

### Live Client Data usage

To safely use Live Client Data API, it's recommended to download the root certificate to validate the HTTPS requests. You can find the certificate on [developer.riotgames.com/docs/lol#game-client-api_root-certificatessl-errors](https://developer.riotgames.com/docs/lol#game-client-api_root-certificatessl-errors)

But if you wish, you can omit the `lol.WithCertificate()` option to use Live Client Data API without the certificate.

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/JoaoDanielRufino/goriot/riot"
	"github.com/JoaoDanielRufino/goriot/riot/lol"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := lol.NewClient(lol.WithCertificate("./riotgames.pem"))
	handleError(err)

	res, err := client.LiveClientData.AllGameData()
	handleError(err)

	pretty, err := json.MarshalIndent(res, "", "  ")
	handleError(err)

	fmt.Println(string(pretty))
}
```

### Data Dragon usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/JoaoDanielRufino/goriot/riot"
	"github.com/JoaoDanielRufino/goriot/riot/lol"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := lol.NewClient()
	handleError(err)

	res, err := client.DataDragon.GetChampions()
	handleError(err)

	pretty, err := json.MarshalIndent(res, "", "  ")
	handleError(err)

	fmt.Println(string(pretty))
}
```
