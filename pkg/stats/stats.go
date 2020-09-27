package stats

import (
	"github.com/ayub94/bank/v2/pkg/types"

)

func Avg(payments []types.Payment) types.Money  {
	
	sum := types.Money(0)
	l := 0
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
		    sum += payment.Amount
	        l++
	    }
	}
	
    
	return sum / (types.Money(l))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money  {
	summincategory := types.Money(0)
	for _, payment := range payments{
		if payment.Category== category {
			if payment.Status != types .StatusFail {
			         summincategory += payment.Amount
			}
		}
		
	}
	return summincategory
}