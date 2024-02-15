package checkout

import "testing"

func TestNewCheckout(t *testing.T) {
  pricingRules := map[string]PricingRule{"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3 }}
  checkout := NewCheckout(pricingRules)

  if len(checkout.scannedItems) != 0 {
    t.Errorf("Checkout has been created with preexisting items of length %d", len(checkout.scannedItems))
  }
}

func TestScanSingleItem(t *testing.T) {
  pricingRules := map[string]PricingRule{"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3 }}

  checkout := NewCheckout(pricingRules)

  checkout.Scan("A")

  expectedTotal := 50
  actualTotal := checkout.GetTotalPrice()

  if actualTotal != expectedTotal {
    t.Errorf("Expected total price %d, but got %d", expectedTotal, actualTotal)
  }
}

func TestScanMultiItem(t *testing.T) {
  pricingRules := map[string]PricingRule{
    "A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3 },
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
