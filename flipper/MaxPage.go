package flipper

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

func GetMaxPage() (page int64) {
	resp, err := http.Get("https://api.hypixel.net/skyblock/auctions")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	page = gjson.Get(string(body), "totalPages").Int()
	return
}
