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
  for k, v := range c.scannedItems {
    totalPrice += calculateItemPrice(c.pricingRules[k], v)
  }
  return totalPrice
}

func calculateItemPrice(pricingRule PricingRule, quantity int) int {
  if pricingRule.SpecialAmount > 0 && quantity >= pricingRule.SpecialAmount {
    numberOfSpecialPriceItemGroups := quantity / pricingRule.SpecialAmount
    numberOfRegularPriceItems := quantity % pricingRule.SpecialAmount

    return (numberOfSpecialPriceItemGroups * pricingRule.SpecialPrice) + (numberOfRegularPriceItems * pricingRule.UnitPrice)
  }

  return quantity * pricingRule.UnitPrice
}
