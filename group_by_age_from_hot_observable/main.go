package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"reactiveX/model"
	"strconv"
)

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(model.Customer{
				ID:        i,
				Name:      fmt.Sprintf("customer %d", i),
				Age:       18,
				TaxNumber: fmt.Sprintf("tax %d", i),
			})
		}

		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(model.Customer{
				ID:        i,
				Name:      fmt.Sprintf("customer %d", i),
				Age:       21,
				TaxNumber: fmt.Sprintf("tax %d", i),
			})
		}

		ch <- rxgo.Of(model.Customer{
			ID:        3,
			Name:      fmt.Sprintf("customer %d", 3),
			Age:       15,
			TaxNumber: fmt.Sprintf("tax %d", 3),
		})

		close(ch)
	}()


	observable := rxgo.FromEventSource(ch).GroupByDynamic(func(item rxgo.Item) string {
		customer := item.V.(model.Customer)
		return strconv.Itoa(customer.Age)
	}, rxgo.WithBufferedChannel(3))

	for i := range observable.Observe() {
		groupedObservable := i.V.(rxgo.GroupedObservable)
		fmt.Printf("New observable: %d\n", groupedObservable.Key)

		for i := range groupedObservable.Observe() {
			fmt.Printf("item: %v\n", i.V)
		}
	}

}
