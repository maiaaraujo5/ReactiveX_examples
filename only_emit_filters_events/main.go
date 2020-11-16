package main

import (
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

		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(model.Customer{
				ID:        3,
				Name:      fmt.Sprintf("customer %d", 3),
				Age:       15,
				TaxNumber: fmt.Sprintf("tax %d", 3),
			})
		}

		close(ch)
	}()

	observable := rxgo.FromChannel(ch).Filter(func(i interface{}) bool {
		customer := i.(model.Customer)
		return customer.Age < 18
	})

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
