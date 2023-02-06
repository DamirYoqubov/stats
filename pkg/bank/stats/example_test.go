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