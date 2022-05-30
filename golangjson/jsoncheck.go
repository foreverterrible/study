package golangjson

import (
	"encoding/json"
	"fmt"
)

type ShopInventory struct {
	Product string `json:"product"`
	Kolvo   int    `json:"kolvo"`
}

func CreateJson() {
	product := &ShopInventory{
		Product: "Молоко",
		Kolvo:   1,
	}

	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
		fmt.Println(product.Kolvo)
	}
}

func CreateJsonBeautify() {
	product := &ShopInventory{
		Product: "Картошка",
		Kolvo:   100,
	}
	data, err := json.MarshalIndent(product, "", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
