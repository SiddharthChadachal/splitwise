package splitwise

func (b *Bill) CalculateSplit() map[string]float64 {
	// result[participantID] = amount
	result := make(map[string]float64)

	// 1. Calculate item-level shares
	for _, item := range b.Items {
		if len(item.SharedBy) == 0 {
			continue
		}

		share := item.Price / float64(len(item.SharedBy))

		for _, pid := range item.SharedBy {
			result[pid] += share
		}
	}

	// 2. Add tax proportionally (if tax exists)
	if b.TaxPercent > 0 {
		subtotal := 0.0
		for _, amount := range result {
			subtotal += amount
		}

		taxAmount := subtotal * (b.TaxPercent / 100.0)

		// proportional tax distribution
		for pid, amount := range result {
			ratio := amount / subtotal
			result[pid] += taxAmount * ratio
		}
	}

	return result
}
