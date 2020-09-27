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
			Status:      types.StatusOk,
		
		},
		{
			ID:           2345,
			Amount:       10_000,
			Status:        types.StatusInProgress,
		},
		{
			ID:           3456,
			Amount:       15_000,
			Status:        types.StatusFail,
		},
		{
			ID:          4567,
			Amount:      15_000,
			Status:      types.StatusOk,
		
		},

	}
		fmt.Println(Avg(payments))


		// Output:
		// 10000

}



func ExampleTotalInCategory()  {

	payments := []types.Payment {

		{
			ID:          1234,
			Amount:      10_000,
			Category:    "Cafe",
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

	} 
	
	fmt.Println(TotalInCategory(payments, "Cafe"))
		// Output:
		// 30000

}