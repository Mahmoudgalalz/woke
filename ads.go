
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type ads struct {
    ad string
    topics  string[]
}

func getAds() err, string{
	res, err := http.Get("https://wokeads.kroking.workers.dev/ads")

	if err != nil {
		return err;
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    return string(responseData)
}