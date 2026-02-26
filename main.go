package main

import (
	wmtracker "WM_tracker/internal"
)

func main() {

	var testData []wmtracker.RequestInput
	testData = append(testData, wmtracker.RequestInput{Slug: "vauban_prime_set", WTS: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "protea_prime_set", WTB: true, WTS: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "wisp_prime_set", WTB: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "arcane_hot_shot", Rank: 5, WTB: true, WTS: true})
	testData = append(testData, wmtracker.RequestInput{Slug: "archon_vitality", Rank: 10, WTS: true})
	wmtracker.PrintMutipleTopOrders(&testData)
}
