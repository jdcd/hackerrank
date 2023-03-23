package min_max

import "errors"

type Item struct {
	ID    int
	Name  string
	Price int
	Desc  string
}

var emptyArrayErr = errors.New("empty items array")

func GetMinMax(myItems []Item) (min int, max int, err error) {
	if len(myItems) == 0 {
		return 0, 0, emptyArrayErr
	}

	min, max = myItems[0].Price, myItems[0].Price
	for _, item := range myItems {
		if item.Price > max {
			max = item.Price
		}

		if item.Price < min {
			min = item.Price
		}
	}

	return
}
