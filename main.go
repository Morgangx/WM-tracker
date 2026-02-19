package main

import (
	wmtracker "WM_tracker/api"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: name, rank(0 = not specified)")
		return
	}
	item_slug, rank := os.Args[1], os.Args[2]
	data, err := wmtracker.GetTopOrders(item_slug, rank)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, order := range data.Data.Sell {
		seller := order.User.IngameName
		price := order.Platinum
		var item_rank int8
		if order.Rank != nil {
			item_rank = *order.Rank
		}

		user_status := order.User.Status
		fmt.Printf("User: %s(%s) - Price: %d - Rank: %o\n", seller, user_status, price, item_rank)
	}

	data2, err := wmtracker.GetOrders(item_slug)
	if err != nil {
		fmt.Println(err)
		return
	}
	top_seller2 := data2.Data[0].User.IngameName
	top_price2 := data2.Data[0].Platinum
	item_rank2 := data2.Data[0].Rank
	fmt.Printf("User: %s\nPrice: %d\nRank: %o\n", top_seller2, top_price2, item_rank2)
}
