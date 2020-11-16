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
				TaxNumber: fmt.Sprintf("tax %d", i),
			})
		}

		close(ch)
	}()

	observable := rxgo.FromEventSource(ch)

	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

}
