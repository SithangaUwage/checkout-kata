package internal_test

import (
	"testing"

	"github.com/SithangaUwage/checkout-kata/internal"
)

func TestIndividualItems(t *testing.T) {
	defaultItems := map[string]internal.Item{
		"A": {UnitPrice: 50, SpecialOffer: internal.SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, SpecialOffer: internal.SpecialOffer{Quantity: 2, Price: 45}},
		"C": {UnitPrice: 20, SpecialOffer: internal.SpecialOffer{}},
		"D": {UnitPrice: 15, SpecialOffer: internal.SpecialOffer{}},
	}

	checkout := internal.InitialiseCheckout(defaultItems)

	checkout.Scan("C")
	checkout.Scan("D")

	totalPrice := checkout.CalculateTotalPrice()
	expectedPrice := defaultItems["C"].UnitPrice + defaultItems["D"].UnitPrice

	if totalPrice != expectedPrice {
		t.Errorf("Expected total price to be %d, got %d", expectedPrice, totalPrice)
	}
}

func TestIndividualItemsSpecialPrice(t *testing.T) {
	defaultItems := map[string]internal.Item{
		"A": {UnitPrice: 50, SpecialOffer: internal.SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, SpecialOffer: internal.SpecialOffer{Quantity: 2, Price: 45}},
		"C": {UnitPrice: 20, SpecialOffer: internal.SpecialOffer{}},
		"D": {UnitPrice: 15, SpecialOffer: internal.SpecialOffer{}},
	}

	checkout := internal.InitialiseCheckout(defaultItems)

	checkout.Scan("A")
	checkout.Scan("A")
	checkout.Scan("A")

	totalPrice := checkout.CalculateTotalPrice()
	expectedPrice := defaultItems["A"].SpecialOffer.Price

	if totalPrice != expectedPrice {
		t.Errorf("Expected total price to be %d, got %d", expectedPrice, totalPrice)
	}
}

func TestItemsInAnyOrder(t *testing.T) {
	defaultItems := map[string]internal.Item{
		"A": {UnitPrice: 50, SpecialOffer: internal.SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, SpecialOffer: internal.SpecialOffer{Quantity: 2, Price: 45}},
		"C": {UnitPrice: 20, SpecialOffer: internal.SpecialOffer{}},
		"D": {UnitPrice: 15, SpecialOffer: internal.SpecialOffer{}},
	}

	checkout := internal.InitialiseCheckout(defaultItems)

	checkout.Scan("B")
	checkout.Scan("A")
	checkout.Scan("D")
	checkout.Scan("A")

	totalPrice := checkout.CalculateTotalPrice()
	expectedPrice := defaultItems["B"].UnitPrice + (defaultItems["A"].UnitPrice * 2) + defaultItems["D"].UnitPrice

	if totalPrice != expectedPrice {
		t.Errorf("Expected total price to be %d, got %d", expectedPrice, totalPrice)
	}
}

func TestUpdateExisitingItem(t *testing.T) {
	defaultItems := map[string]internal.Item{
		"A": {UnitPrice: 50, SpecialOffer: internal.SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, SpecialOffer: internal.SpecialOffer{Quantity: 2, Price: 45}},
		"C": {UnitPrice: 20, SpecialOffer: internal.SpecialOffer{}},
		"D": {UnitPrice: 15, SpecialOffer: internal.SpecialOffer{}},
	}

	internal.UpdateItems(defaultItems, map[string]internal.Item{
		"A": {UnitPrice: 60, SpecialOffer: internal.SpecialOffer{}},
	})

	if defaultItems["A"].UnitPrice != 60 {
		t.Errorf("Expected item to be %v, got %v", 60, defaultItems["A"])
	}
}

func TestAddingNewItem(t *testing.T) {
	defaultItems := map[string]internal.Item{
		"A": {UnitPrice: 50, SpecialOffer: internal.SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, SpecialOffer: internal.SpecialOffer{Quantity: 2, Price: 45}},
		"C": {UnitPrice: 20, SpecialOffer: internal.SpecialOffer{}},
		"D": {UnitPrice: 15, SpecialOffer: internal.SpecialOffer{}},
	}

	internal.UpdateItems(defaultItems, map[string]internal.Item{
		"E": {UnitPrice: 10, SpecialOffer: internal.SpecialOffer{Quantity: 2, Price: 5}},
	})

	if _, exists := defaultItems["E"]; !exists {
		t.Errorf("Expected item to exist")
	}
}
