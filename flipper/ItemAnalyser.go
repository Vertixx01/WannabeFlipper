// ✪
package flipper

import (
	"strings"
)

type AnalysedData struct {
	Stars []struct {
		Five  int
		Four  int
		Three int
		Two   int
		One   int
	}
}

func ItemAnalyser() (data AnalysedData) {
	data.Stars = append(data.Stars, struct {
		Five  int
		Four  int
		Three int
		Two   int
		One   int
	}{
		Five:  0,
		Four:  0,
		Three: 0,
		Two:   0,
		One:   0,
	})
	for _, item := range ItemCounter().ItemCount {
		switch strings.Count(item.Name, "✪") {
		case 5:
			data.Stars[0].Five++
		case 4:
			data.Stars[0].Four++
		case 3:
			data.Stars[0].Three++
		case 2:
			data.Stars[0].Two++
		case 1:
			data.Stars[0].One++
		}
	}
	return
}
