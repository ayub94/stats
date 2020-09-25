package stats

import (
	"fmt"
	"github.com/ayub94/bank/pkg/types"
)

func ExampleAvg()  {

	payments := []types.Payment {

		{
			ID:          1234,
			Amount:      5000,
			Category:    "Cat",
		},
		{
			ID:           2345,
			Amount:       10_000,
			Category:     "Cat",
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Cat",
		},

	}
	result := Avg(payments)
		fmt.Println(result)


		// Output:
		// 10000

}



func ExampleTotalInCategory()  {

	payments := []types.Payment {

		{
			ID:          1234,
			Amount:      5000,
			Category:    "Shopping",
		},
		{
			ID:           2345,
			Amount:       10_000,
			Category:     "Cafe",
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Restaurant",
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Cafe",
		},
	} 
	
	fmt.Println(TotalInCategory(payments, "Cafe"))
		// Output:
		// 25000

}