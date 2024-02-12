package internal

import "fmt"

type Item struct {
	UnitPrice    int64
	SpecialOffer SpecialOffer
}

type SpecialOffer struct {
	Quantity int64
	Price    int64
}

type Checkout struct {
	Items        map[string]Item
	ScannedItems map[string]int64
}

func defaultItems() map[string]Item {
	return map[string]Item{
		"A": {UnitPrice: 50, SpecialOffer: SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, SpecialOffer: SpecialOffer{Quantity: 2, Price: 45}},
		"C": {UnitPrice: 20, SpecialOffer: SpecialOffer{}},
		"D": {UnitPrice: 15, SpecialOffer: SpecialOffer{}},
	}
}

func InitialiseCheckout(item map[string]Item) Checkout {
	return Checkout{
		Items:        item,
		ScannedItems: make(map[string]int64),
	}
}

func (c *Checkout) Scan(item string) {
	if _, exists := c.Items[item]; exists {
		c.ScannedItems[item]++
	}
}

func (c *Checkout) calculateItemPrice(item string, quantity int64) int64 {
	var itemPrice int64 = 0

	if c.Items[item].SpecialOffer.Quantity != 0 && quantity >= c.Items[item].SpecialOffer.Quantity {
		eligibleForSpecialOffer := quantity / c.Items[item].SpecialOffer.Quantity
		nonEligibleItems := quantity % c.Items[item].SpecialOffer.Quantity

		itemPrice += eligibleForSpecialOffer * c.Items[item].SpecialOffer.Price
		itemPrice += nonEligibleItems * c.Items[item].UnitPrice
	} else {
		itemPrice += quantity * c.Items[item].UnitPrice
	}

	return itemPrice
}

func (c *Checkout) CalculateTotalPrice() int64 {
	var totalPrice int64 = 0

	for item, quantity := range c.ScannedItems {
		totalPrice += c.calculateItemPrice(item, quantity)
	}
	return totalPrice
}

func UpdateItems(items map[string]Item, newItems map[string]Item) {
	for key, newItem := range newItems {
		items[key] = newItem
	}
}

func StoreCheckout() {
	items := defaultItems()
	checkout := InitialiseCheckout(items)

	// Add items and calculate total price
	checkout.Scan("A")
	checkout.Scan("C")
	checkout.Scan("B")
	checkout.Scan("B")
	checkout.Scan("D")
	checkout.Scan("A")
	checkout.Scan("A")

	fmt.Printf("Total price: %+v\n", checkout.CalculateTotalPrice())
}

func StoreUpdateitems() {
	items := defaultItems()

	// Update Item A and add new Item E
	UpdateItems(items, map[string]Item{
		"A": {UnitPrice: 60, SpecialOffer: SpecialOffer{}},
		"E": {UnitPrice: 10, SpecialOffer: SpecialOffer{
			Quantity: 5, Price: 30,
		}},
	})

	for key, value := range items {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}
}
