package stats

import (
	"github.com/ayub94/bank/v2/pkg/types"

)

func Avg(payments []types.Payment) types.Money  {
	
	averagesum := types.Money(0)
    sum := types.Money(0)
	for _, payment := range payments {
		if payment.Status == "Ok" || payment.Status == "INPROGRESS"{
	            sum += payment.Amount
	            averagesum = sum / (types.Money(len(payments)))
		}
	
	}

	return averagesum
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	summincategory := types.Money(0)
	for _, payment := range payments{
		if payment.Category== category {
			if payment.Status == "Ok" || payment.Status == "INPROGRESS" {
			summincategory += payment.Amount
			}
		}
		
	}
	return summincategory
}