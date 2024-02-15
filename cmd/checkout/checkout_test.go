package checkout

import "testing"

func TestNewCheckout(t *testing.T) {
	pricingRules := map[string]PricingRule{"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3}}
	checkout := NewCheckout(pricingRules)

	if len(checkout.scannedItems) != 0 {
		t.Errorf("Checkout has been created with preexisting items of length %d", len(checkout.scannedItems))
	}
}

func TestScanSingleItem(t *testing.T) {
	pricingRules := map[string]PricingRule{"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3}}

	checkout := NewCheckout(pricingRules)

	checkout.Scan("A")

	expectedTotal := 50
	actualTotal := checkout.GetTotalPrice()

	if actualTotal != expectedTotal {
		t.Errorf("Expected total price %d, but got %d", expectedTotal, actualTotal)
	}
}

func TestRemoveSingleItem(t *testing.T) {
	pricingRules := map[string]PricingRule{"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3}}

	checkout := NewCheckout(pricingRules)

	checkout.Scan("A")
	checkout.Scan("A")
	checkout.Scan("A")
	checkout.RemoveItem("A")

	expectedTotalNumItems := 2
	expectedTotalPrice := 100

	actualNumOfItems := checkout.scannedItems["A"]
	actualTotalPrice := checkout.GetTotalPrice()

	if actualNumOfItems != expectedTotalNumItems {
		t.Errorf("Expected total number of items to be %d, but got %d", expectedTotalNumItems, actualNumOfItems)
	}

	if actualTotalPrice != expectedTotalPrice {
		t.Errorf("Expected total price %d, but got %d", expectedTotalPrice, actualTotalPrice)
	}
}

func TestScanMultiItem(t *testing.T) {
	pricingRules := map[string]PricingRule{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3},
	}

	checkout := NewCheckout(pricingRules)

	checkout.Scan("A")
	checkout.Scan("A")
	checkout.Scan("A")

	expectedTotal := 130
	actualTotal := checkout.GetTotalPrice()

	if actualTotal != expectedTotal {
		t.Errorf("Expected total price %d, but got %d", expectedTotal, actualTotal)
	}
}

func TestScanMultiItemTypes(t *testing.T) {
	pricingRules := map[string]PricingRule{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialAmount: 2},
		"C": {UnitPrice: 20},
	}

	checkout := NewCheckout(pricingRules)

	checkout.Scan("A")
	checkout.Scan("B")
	checkout.Scan("C")
	checkout.Scan("A")
	checkout.Scan("A")

	expectedTotal := 180
	actualTotal := checkout.GetTotalPrice()

	if actualTotal != expectedTotal {
		t.Errorf("Expected total price %d, but got %d", expectedTotal, actualTotal)
	}
}

func TestPriceModification(t *testing.T) {
	pricingRules := map[string]PricingRule{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialAmount: 2},
		"C": {UnitPrice: 20},
	}

	newPricingRules := map[string]PricingRule{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialAmount: 2},
		"C": {UnitPrice: 2020},
	}

	checkout := NewCheckout(pricingRules)

	checkout.Scan("A")
	checkout.Scan("B")
	checkout.Scan("C")
	checkout.Scan("A")
	checkout.Scan("A")

	checkout.ModifyPricingRules(newPricingRules)

	expectedTotal := 2180
	actualTotal := checkout.GetTotalPrice()

	if actualTotal != expectedTotal {
		t.Errorf("Expected total price %d, but got %d", expectedTotal, actualTotal)
	}
}
