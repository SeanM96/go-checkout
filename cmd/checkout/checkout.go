package checkout

import "log"

type PricingRule struct {
	UnitPrice     int
	SpecialPrice  int
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

func (c *Checkout) ModifyPricingRules(newPricingRules map[string]PricingRule) {
	c.pricingRules = newPricingRules
}

func (c *Checkout) Scan(item string) {
	if _, ok := c.pricingRules[item]; ok {
		c.scannedItems[item]++
	} else {
		log.Printf("Warning:%s %s", item, "is an unrecognised item!")
	}
}

func (c *Checkout) RemoveItem(item string) {
	if count, ok := c.scannedItems[item]; ok && count > 0 {
		c.scannedItems[item]--
	} else {
		log.Printf("Warning: %s %s", item, "cannot be removed as it has not been scanned.")
	}
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
