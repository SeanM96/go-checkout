package main

import (
	"log"

	"github.com/SeanM96/thinkmoney-checkout/cmd/checkout/checkout"
)
func main() {
  pricingRules := map[string]checkout.PricingRule{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialAmount: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialAmount: 2},
		"C": {UnitPrice: 20},
	}

  newPricingRules := map[string]checkout.PricingRule{
		"A": {UnitPrice: 500, SpecialPrice: 1000, SpecialAmount: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialAmount: 2},
		"C": {UnitPrice: 20},
	}

  checkout := checkout.NewCheckout(pricingRules)

  checkout.Scan("A")
  checkout.Scan("A") // Total of A, 2

  checkout.Scan("B") // Total of B, 1

  checkout.Scan("A") // Total of A 3, SpecialPrice kicks in
  totalPrice := checkout.GetTotalPrice()
  log.Printf("Total price is: %d", totalPrice)

  checkout.Scan("A") // Total of A 4, regular price resumes
  totalPrice = checkout.GetTotalPrice()
  log.Printf("Total price is: %d", totalPrice)

  checkout.Scan("C") // Total of C 1
  checkout.RemoveItem("A") // Remove Item A. Special price still resumes for the first 3

  totalPrice = checkout.GetTotalPrice()
  log.Printf("Total price is: %d", totalPrice)

  checkout.ModifyPricingRules(newPricingRules)
  totalPrice = checkout.GetTotalPrice()
  log.Printf("Total price is: %d", totalPrice)
}
