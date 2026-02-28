package wmtracker

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func MainLoop(refresh *int, loop *bool) {
	for {
		data, err := LoadRequests()
		if err != nil {
			fmt.Println("Error loading requests:", err)
			return
		}
		fmt.Println("Starting main loop iteration")
		PrintMutipleTopOrders(&data)
		time.Sleep(time.Duration(*refresh) * time.Second)
		if !*loop {
			break
		}
	}
}

func PrintMutipleTopOrders(requests *[]RequestInput) {
	for i, req := range *requests {
		if i%3 == 0 { // API has a limit of 3 requests per second
			time.Sleep(1200 * time.Millisecond)
		}
		data, err := GetTopOrders(req.Slug, req.Rank)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !req.WTS && !req.WTB {
			continue
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "----%d---- \t%s\n", i+1, req.Slug)
		if req.WTS {
			if len(data.Data.Sell) == 0 {
				continue
			}
			fmt.Fprintln(w, "WTS ORDERS")
			fmt.Fprintln(w, "i\tSELLER\tSTATUS\tPRICE\tRANK")
			for j, order := range data.Data.Sell {
				seller := order.User.IngameName
				price := order.Platinum
				var item_rank int8
				if order.Rank != nil {
					item_rank = *order.Rank
				}

				user_status := order.User.Status
				fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%o\n", j+1, seller, user_status, price, item_rank)
			}
		}
		if req.WTB {
			if len(data.Data.Buy) == 0 {
				continue
			}
			fmt.Fprintln(w, "WTB ORDERS")
			fmt.Fprintln(w, "i\tBUYER\tSTATUS\tPRICE\tRANK")
			for j, order := range data.Data.Buy {
				seller := order.User.IngameName
				price := order.Platinum
				var item_rank int8
				if order.Rank != nil {
					item_rank = *order.Rank
				}

				user_status := order.User.Status
				fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%o\n", j+1, seller, user_status, price, item_rank)
			}
		}
		w.Flush()
	}
}
