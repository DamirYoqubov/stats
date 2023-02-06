package stats

import "github.com/DamirYoqubov/stats/pkg/bank/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	
	sum := types.Money(0)
	
	count := int32(0)

	for _, payment := range payments {
		sum += payment.Amount
		count++
	}
	return sum/types.Money(count)
}