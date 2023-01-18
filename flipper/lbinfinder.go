package flipper

import (
	"fmt"
	"os"
	"strings"

	"github.com/Vertixx01/wannabeflipper/utils"
)

type LbinData struct {
	Lbin []struct {
		UUID   string
		Name   string
		Price  string
		TIER   string
		Lore   string
		Margin string
		Profit string
	}
}

type TempItem struct {
	TempItem []struct {
		UUID  string
		Name  string
		Price string
		TIER  string
		Lore  string
	}
}

func LbinFinder() (data LbinData) {
	auctions := Looper()
	var tempItem TempItem
	config, _ := os.Open("config.json")
	defer config.Close()
	// minMargin := gjson.Get(utils.Read(config), "minMargin").Float()
	// minProfit := gjson.Get(utils.Read(config), "minProfit").Float()
	for _, auction := range auctions {
		for _, item := range auction.Auctions {
			if strings.Contains(item.UUID, tempItem.TempItem[0].UUID) {
				utils.NewLogger().Info.Println(fmt.Sprintf("%s Found new item: %s", utils.TimeStamp(), item.Name))
				tempItem.TempItem = append(tempItem.TempItem, struct {
					UUID  string
					Name  string
					Price string
					TIER  string
					Lore  string
				}{
					UUID:  item.UUID,
					Name:  item.Name,
					Price: item.Price,
					TIER:  item.TIER,
					Lore:  item.Lore,
				})
			}
		}
	}
	return data
}
