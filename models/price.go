package models

import (
	"errors"
	"strconv"
	"time"
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

func AddOne(Price Price) (PriceId string) {
	Price.PriceId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Prices[Price.PriceId] = &Price
	return Price.PriceId
}

func GetOne(PriceId string) (Price *Price, err error) {
	if v, ok := Prices[PriceId]; ok {
		return v, nil
	}
	return nil, errors.New("PriceId Not Exist")
}

func GetAll() map[string]*Price {
	return Prices
}

func Update(PriceId string, Score int64) (err error) {
	if v, ok := Prices[PriceId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("PriceId Not Exist")
}

func Delete(PriceId string) {
	delete(Prices, PriceId)
}
