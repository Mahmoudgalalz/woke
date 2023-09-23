package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type Ad struct {
	Ad     string   `json:"ad"`
	Topics []string `json:"topics"`
	Link   string   `json:"link"`
}

func GetAds() ([]Ad, error) {
	res, err := http.Get("https://wokeads.kroking.workers.dev/ads")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var ads []Ad
	err = json.Unmarshal(responseData, &ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}
func Randomize(ads []Ad) Ad {
	limit := len(ads)
	randomIdx := rand.Intn(limit)

	return ads[randomIdx]
}

func RandomAd() (Ad, error) {
	ads, err := GetAds()
	if err != nil {
		return Ad{}, err
	}

	if len(ads) == 0 {
		return Ad{}, errors.New("no ads found")
	}

	ad := Randomize(ads)
	return ad, nil
}
