package wmtracker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl string = "https://api.warframe.market/v2/"

func GetTopOrders(item_slug string, rank int8) (*TopOrdersResponse, error) {
	rankSuff := ""
	if rank != 0 {
		rankSuff = fmt.Sprintf("?rank=%d", rank)
	}
	fullUrl := baseUrl + "orders/item/" + item_slug + "/top" + rankSuff
	res, err := http.Get(fullUrl)
	if err != nil {
		return nil, fmt.Errorf("Could not get response: %w", err)
	}
	defer res.Body.Close()

	var data TopOrdersResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("Error decoding response: %w", err)
	}
	return &data, nil
}

func PrintMutipleTopOrders(requests *[]RequestData) {
	for i, req := range *requests {
		data, err := GetTopOrders(req.Slug, req.Rank)
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
			fmt.Printf("%d.) User: %s(%s) - Price: %d - Rank: %o\n", i+1, seller, user_status, price, item_rank)
		}
	}

}

func GetOrders(item_slug string) (*OrdersResponse, error) {
	fullUrl := baseUrl + "orders/item/" + item_slug
	res, err := http.Get(fullUrl)
	if err != nil {
		return nil, fmt.Errorf("Could not get response: %w", err)
	}
	defer res.Body.Close()
	var data OrdersResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("Error decoding response: %w", err)
	}
	return &data, nil
}
