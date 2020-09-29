package stats

import (
	"github.com/ayub94/bank/v2/pkg/types"

)

// FilteredByCategory возврашает платежи в указаной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment{

	var filtered []types.Payment
	
	for _, payment := range payments {
	  if payment.Category == category {
		  filtered = append(filtered, payment)
	  }
		
	}

return filtered

}
// CategoryAvg counts the avarage sum of categories
func CategoriesAvg(payments []types.Payment) map [types.Category]types.Money  {
  categories := map [types.Category] types.Money {}
  count := map [types.Category] types.Money {}

	  for _, payment := range payments{
		  if payment.Status != types.StatusFail {
				  categories [payment.Category] += payment.Amount
				  count [payment.Category] ++		  
		  }
		for key := range categories {
			categories[key] /= count[key]
		}  
	  }	
		return categories
 }