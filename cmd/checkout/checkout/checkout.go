package checkout

type PricingRule struct {
  UnitPrice float32
  SpecialPrice float32
  SpecialAmount int8
}

type Checkout struct {
pricingRules map[string]PricingRule
scannedItems map[string]int
}

func NewCheckout(pricingRules map[string]PricingRule) *Checkout {
  return &Checkout{
    pricingRules: pricingRules,
    scannedItems: make(map[string]int),
  }
}
