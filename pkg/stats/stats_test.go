package stats

import(
	"github.com/ayub94/bank/v2/pkg/types"
	"testing"
	"reflect"
)

func TestFilterByCategory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")
	 if len(result) !=0 {
		 	t.Error("len(result) !=0")
	 }
}
func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")
	 if len(result) !=0 {
		 	t.Error("len(result) !=0")
	 }
}
func TestFilterByCategory_NotFound(t *testing.T) {
		payments := []types.Payment{
			{ID: 1,  Category: "auto"},
			{ID: 2,  Category: "food"},
			{ID: 3,  Category: "auto"},
			{ID: 4,  Category: "fun"},
			{ID: 5,  Category: "auto"},
			{ID: 6,  Category: "auto"},
			{ID: 7,  Category: "auto"},
		}
		result := FilterByCategory(payments, "mobile")
		 if len(result) !=0 {
				 t.Error("len(result) !=0")
		 }
}
func TestFilterByCategory_FoundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1,  Category: "auto"},
		{ID: 2,  Category: "food"},
		{ID: 3,  Category: "auto"},
		{ID: 4,  Category: "fun"},
		{ID: 5,  Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 4, Category: "fun"},
	}	
	result := FilterByCategory(payments, "fun")
	 if !reflect.DeepEqual(expected, result) {
			 t.Error("Invalid result")
	 }
}
func TestFilterByCategory_FoundSeveral(t *testing.T) {
	payments := []types.Payment{
		{ID: 1,  Category: "auto"},
		{ID: 2,  Category: "food"},
		{ID: 3,  Category: "auto"},
		{ID: 4,  Category: "fun"},
		{ID: 5,  Category: "auto"},
	}
	expected := []types.Payment{
		{ID: 1,  Category: "auto"},
		{ID: 3,  Category: "auto"},
		{ID: 5,  Category: "auto"},
	}	
	result := FilterByCategory(payments, "auto")
	 if !reflect.DeepEqual(expected, result) {
			 t.Errorf("Invalid result, expected: %v, actual: %v", expected, result)

	 }
}
 
func TestCategoriesAvgUser(t *testing.T) {
	payments := []types.Payment {
		{ID:1, Category: "auto", Amount: 1_000},
		{ID:2, Category: "mobile", Amount: 2_000},
		{ID:3, Category: "fun", Amount: 3_000},
		{ID:4, Category: "auto", Amount: 4_000},
		{ID:5, Category: "fun", Amount: 5_000},
		{ID:6, Category: "fun", Amount: 5_000},
		{ID:7, Category: "mobile", Amount: 4_000},

	}
	expected := map [types.Category]types.Money{
		"auto":  2_500,
		"mobile":  3_000,
		"fun":  4333,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result){
		t.Errorf("Invalid result, expected: %v, actual: %v\n", expected, result)
	}


}
