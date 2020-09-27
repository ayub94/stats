package stats

import (
	"github.com/ayub94/bank/v2/pkg/types"

)

func Avg(payments []types.Payment) types.Money  {
	
	averagesum := types.Money(0)
	sum := types.Money(0)
	l := 0
	for _, payment := range payments {
		if payment.Status == types.StatusOk || payment.Status == types.StatusInProgress {

			    l+=1
	            sum += payment.Amount
		}
		
		averagesum = sum / (types.Money(l))
	}

	return averagesum
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	summincategory := types.Money(0)
	for _, payment := range payments{
		if payment.Category== category {
			if payment.Status == types.StatusOk || payment.Status == types.StatusInProgress {
			         summincategory += payment.Amount
			}
		}
		
	}
	return summincategory
}