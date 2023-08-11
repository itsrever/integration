package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSunglassesLineItem(t *testing.T) {
	li := sunglassesLineItem()
	assert.Equal(t, li.UnitPrice.AmountCustomer.Amount, sunglassesUnitPrice)
	assert.Equal(t, li.TotalDiscounts.AmountCustomer.Amount, 2.59)
	assert.Equal(t, li.Total.AmountCustomer.Amount,
		li.Subtotal.AmountCustomer.Amount-
			li.TotalDiscounts.AmountCustomer.Amount+
			li.TotalTaxes.AmountCustomer.Amount)
	assert.Greater(t, li.Total.AmountCustomer.Amount, float64(0))
}
