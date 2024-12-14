package lab7

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	ApplyDiscount(discount float64)
}

type Item struct {
	Name  string
	Price float64
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) GetPrice() float64 {
	return i.Price
}

func (i *Item) ApplyDiscount(discount float64) {
	i.Price -= i.Price * discount / 100
}

func CalculateTotal(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func Executelab7() {
	item1 := &Item{Name: "Laptop", Price: 1000}
	item2 := &Item{Name: "Phone", Price: 500}
	item3 := &Item{Name: "Tablet", Price: 300}

	products := []Product{item1, item2, item3}

	fmt.Printf("Total before discount: %.2f\n", CalculateTotal(products))

	for _, product := range products {
		product.ApplyDiscount(10)
	}

	fmt.Printf("Total after discount: %.2f\n", CalculateTotal(products))
}
