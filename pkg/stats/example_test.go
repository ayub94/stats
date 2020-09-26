package stats

import (
	"fmt"
	"github.com/ayub94/bank/v2/pkg/types"
)

func ExampleAvg()  {

	payments := []types.Payment {

		{
			ID:          1234,
			Amount:      5000,
			Category:    "Cat",
			Status:      "Ok",
		
		},
		{
			ID:           2345,
			Amount:       10_000,
			Category:     "Cat",
			Status:        "INPROGRESS",
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Cat",
			Status:        "FAIL",
		},

	}
	result := Avg(payments)
		fmt.Println(result)


		// Output:
		// 15000

}



func ExampleTotalInCategory()  {

	payments := []types.Payment {

		{
			ID:          1234,
			Amount:      5000,
			Category:    "Shopping",
			Status:       "Ok",
		},
		{
			ID:           2345,
			Amount:       10_000,
			Category:     "Cafe",
			Status:         "FAIL",
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Restaurant",
			Status:        "INPROGRESS",  
		},
		{
			ID:           3456,
			Amount:       15_000,
			Category:      "Cafe",
			Status:        "Ok",
		},
	} 
	
	fmt.Println(TotalInCategory(payments, "Cafe"))
		// Output:
		// 35000

}