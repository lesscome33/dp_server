package models

import (
)

var (
	Prices map[string]*Price
)

type Price struct {
	PriceId    string
	Score      int64
	PlayerName string
}

func init() {
	Prices = make(map[string]*Price)
	Prices["hjkhsbnmn123"] = &Price{"hjkhsbnmn123", 100, "astaxie"}
	Prices["mjjkxsxsaa23"] = &Price{"mjjkxsxsaa23", 101, "someone"}
}
