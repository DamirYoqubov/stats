package stats

import (
	"fmt"

	"github.com/DamirYoqubov/stats/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 3,
		},
		{
			Amount: 4,
		},
		{
			Amount: 5,
		},
	}
	avg := Avg(payments)
	fmt.Println(avg)
	//Output: 4
}

func ExampleTotalInCategory()  {
	category := types.Category("Auto")
	category2 := types.Category("car")
	payments := []types.Payment{
		{
			Category: category,
			Amount: 50_00,
		},
		{
			Category: category,
			Amount: 10_00,
		},
		{
			Category: category2,
			Amount: 20_00,
		},
	}
	result := TotalInCategory(payments, category)

	fmt.Println(result)
	//Output: 6000
}