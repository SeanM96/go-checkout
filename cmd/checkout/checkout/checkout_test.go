package checkout

import "testing"

func TestNewCheckout(t *testing.T) {
  pricingRules := map[string]PricingRule{"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3 }}
  checkout := NewCheckout(pricingRules)

  if len(checkout.scannedItems) != 0 {
    t.Errorf("Checkout has been created with preexisting items of length %d", len(checkout.scannedItems))
  }
}
