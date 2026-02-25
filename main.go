package main

import (
	wmtracker "WM_tracker/api"
)

func main() {

	var testData []wmtracker.RequestData
	r1 := wmtracker.RequestData{Slug: "vauban_prime_set", Rank: 0}
	r2 := wmtracker.RequestData{Slug: "protea_prime_set", Rank: 0}
	testData = append(testData, r1)
	testData = append(testData, r2)
	wmtracker.PrintMutipleTopOrders(&testData)
}
