package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"reactiveX/model"
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
		for i := 0; i < 2; i++ {
			ch <- rxgo.Of(model.Customer{
				ID:        i,
				Name:      fmt.Sprintf("customer %d", i),
				Age:       18,
				TaxNumber: fmt.Sprintf("tax %d", i),
			})
		}
		close(ch)
	}()

	observable := rxgo.FromChannel(ch).Distinct(func(ctx context.Context, i interface{}) (interface{}, error) {
		customer := i.(model.Customer)
		return customer.Age, nil
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
