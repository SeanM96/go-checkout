# Checkout System

Checkout System is a Go implementation of a simple checkout system. It supports multiple pricing schemes, such as special discounts for purchasing multiple items of a particular type.

## Features

- Scan items to add them to the checkout.
- Remove scanned items from the checkout.
- Modify pricing rules dynamically during the checkout process.
- Calculate the total price based on the scanned items and pricing rules.

## Getting Started

### Prerequisites

- Go
- Git

### Installation

```bash
git clone https://github.com/SeanM96/go-checkout.git
cd go-checkout/cmd/checkout
go run main.go
```

### Usage

1. Import the checkout package in your Go project.
2. Create a new checkout instance with pricing rules.
3. Scan items using the Scan method.
4. Optionally remove scanned items using the RemoveItem method.
5. Optionally modify pricing rules using the ModifyPricingRules method.
6. Get the total price using the GetTotalPrice method.

An example is shown in `main.go`


### Known Improvements

1. Curently, if you modify the pricing rules during a checkout, it will retroactively change the prices. It's unclear if this would be desired functionality, so no special behaviour has been implemented. If desired, the scanned items map could contain item info, such as "priceAtTimeOfScan", instead of just the count.
2. Linked to the previous, there is no finish sale method. This would be fairly easy to implement, in a basic k-v store.
3. No interface. I believe in not optimising too early - code should be designed to be easily extensible, so should be written with possible future requirements in mind, but not tied to them. (Johnny Boursiquot put it far better than I can in a recent GoTime episode https://changelog.com/gotime/294#transcript-316).

