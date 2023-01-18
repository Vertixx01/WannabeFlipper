package flipper

import (
	"io/ioutil"
	"net/http"

	"github.com/Vertixx01/wannabeflipper/utils"
	"github.com/tidwall/gjson"
)

type Auctions struct {
	Auctions []struct {
		UUID  string `json:"uuid"`
		Name  string `json:"item_name"`
		Price string `json:"starting_bid"`
		TIER  string `json:"tier"`
		Lore  string `json:"item_lore"`
	} `json:"auctions"`
}

func GetPage(page string) (AuctionsList []Auctions) {
	var data Auctions
	resp, err := http.Get("https://api.hypixel.net/skyblock/auctions?page=" + page)
	if err != nil {
		utils.NewLogger().Error.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	auctions := gjson.Get(string(body), "auctions").Array()
	for _, auction := range auctions {
		if auction.Get("bin").Bool() {
			data.Auctions = append(data.Auctions, struct {
				UUID  string `json:"uuid"`
				Name  string `json:"item_name"`
				Price string `json:"starting_bid"`
				TIER  string `json:"tier"`
				Lore  string `json:"item_lore"`
			}{
				UUID:  auction.Get("uuid").String(),
				Name:  auction.Get("item_name").String(),
				Price: auction.Get("starting_bid").String(),
				TIER:  auction.Get("tier").String(),
				Lore:  auction.Get("item_lore").String(),
			})
		}
	}
	AuctionsList = append(AuctionsList, data)
	return
}
