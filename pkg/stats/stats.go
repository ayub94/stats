package stats

import (
	"github.com/ayub94/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money  {
	
	averagesum := types.Money(0)
    sum := types.Money(0)
	for _, payment := range payments {
	sum += payment.Amount
	averagesum = sum / (types.Money(len(payments)))
	
	}

	return averagesum
}
//const countcategory = "Cafe"
func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	summincategory := types.Money(0)
	for _, payment := range payments{
		if payment.Category== category {
			summincategory += payment.Amount
		}
		
	}
	return summincategory
}