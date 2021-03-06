package stats

import (
	"github.com/ayub94/bank/v2/pkg/types"

)
//Avg func return average amount from slice Payment
func Avg(payments []types.Payment) types.Money {
  var sum types.Money
  i := 0
  for _, v := range payments {
    if v.Status == types.StatusFail {
      continue
    }
    sum += v.Amount
    i++
  }
  return sum / types.Money(i)
}
//CategoriesAvg func
func CategoriesAvg(payments []types.Payment) ( map[types.Category]types.Money) {

  mp := make(map[types.Category]types.Money)
  for _, v := range payments {

    if _, er := mp[v.Category]; er {
      continue
    }
    filtered := FilterByCategory(payments, v.Category)
    mp[v.Category]=Avg(filtered)
  }


  return mp
}


//FilterByCategory func
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment{

  var filtered []types.Payment

  for _, v := range payments {
    if v.Category == category {
      filtered = append(filtered, v)
    }
  }
  return filtered
}