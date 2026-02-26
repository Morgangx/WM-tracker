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
