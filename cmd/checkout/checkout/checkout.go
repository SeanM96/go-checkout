package checkout

type PricingRule struct {
  UnitPrice int
  SpecialPrice int
  SpecialAmount int
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

func (c *Checkout) Scan(item string) {
  c.scannedItems[item]++
}

func (c *Checkout) GetTotalPrice() int {
  totalPrice := 0
  for k := range c.scannedItems {
    totalPrice += c.pricingRules[k].UnitPrice
  }
  return totalPrice
}
