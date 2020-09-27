package stats

import (
	"fmt"
	"github.com/ayub94/bank/v2/pkg/types"
)

func ExampleAvg()  {

	payments := []types.Payment {

		{
			ID:          1234,
			Amount:      5_000,
			Category:    "Cat",
			Status:      types.StatusOk,
		
		},
		{
			ID:           2345,
			Amount:       10_000,
			Category:     "Cat",
			Status:        types.StatusInProgress,
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Cat",
			Status:        types.StatusFail,
		},
		{
			ID:          4567,
			Amount:      -5_000,
			Category:    "Cat",
			Status:      types.StatusOk,
		
		},
		{
			ID:          4567,
			Amount:      15_000,
			Category:    "Cat",
			Status:      types.StatusOk,
		
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
			Amount:      5_000,
			Category:    "Shopping",
			Status:       types.StatusOk,
		},
		{
			ID:           2345,
			Amount:       10_000,
			Category:     "Cafe",
			Status:       types.StatusFail,
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Restaurant",
			Status:        types.StatusInProgress,  
		},
		{
			ID:           3456,
			Amount:       20_000,
			Category:      "Cafe",
			Status:        types.StatusOk,
		},
		{
			ID:          4567,
			Amount:      -5_000,
			Category:    "Cafe",
			Status:      types.StatusInProgress,
		
		},
	} 
	
	fmt.Println(TotalInCategory(payments, "Cafe"))
		// Output:
		// 20000

}