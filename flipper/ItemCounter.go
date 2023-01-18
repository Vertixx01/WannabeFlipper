package flipper

type ItemCount struct {
	ItemCount []struct {
		Name  string
		Count int
	}
}

func ItemCounter() (data ItemCount) {
	auctions := Looper()
	for _, auction := range auctions[0].Auctions {
		var found bool
		for index, item := range data.ItemCount {
			if item.Name == auction.Name {
				data.ItemCount[index].Count++
				found = true
				break
			}
		}
		if !found {
			data.ItemCount = append(data.ItemCount, struct {
				Name  string
				Count int
			}{
				Name:  auction.Name,
				Count: 1,
			})
		}
	}
	return
}
