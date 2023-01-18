package flipper

import (
	"fmt"
)

func AuctionSender(name string, price1 string, price2 string, profit string) string {
	msg := fmt.Sprintf("%s %s -> %s (+%s)", name, price1, price2, profit)
	return msg
}
