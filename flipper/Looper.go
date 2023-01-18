package flipper

import "fmt"

type AuctionList struct {
	Auctions []Auctions
}

func Looper() (data []Auctions) {
	//page := GetMaxPage()
	for i := 0; int64(i) <= 3; i++ {
		auction := GetPage(fmt.Sprintf("%v", i))
		data = append(data, auction...)
	}
	return data
}
