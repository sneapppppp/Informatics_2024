package main

import (
	"fmt"
)

type Product interface {
	GetName() string
	GetPrice() float64
	SetPrice(price float64)
	ApplyDiscount(discountPercentage float64)
}

type BaseProduct struct {
	Name  string
	Price float64
}

func (p *BaseProduct) GetName() string {
	return p.Name
}

func (p *BaseProduct) GetPrice() float64 {
	return p.Price
}

func (p *BaseProduct) SetPrice(price float64) {
	p.Price = price
}

func (p *BaseProduct) ApplyDiscount(discountPercentage float64) {
	p.Price = p.Price * (1 - discountPercentage/100)
}

type Clothing struct {
	BaseProduct
	Size  string
	Color string
}

func (c *Clothing) UpdateCharacteristics(size, color string) {
	c.Size = size
	c.Color = color
}

type Electronics struct {
	BaseProduct
	Warranty int 
}

func (e *Electronics) UpdateWarranty(warranty int) {
	e.Warranty = warranty
}

func CalculateTotalPrice(products []Product) float64 {
	var total float64
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func main() {
	clothing := &Clothing{
		BaseProduct: BaseProduct{Name: "T-Shirt", Price: 20.0},
		Size:        "M",
		Color:       "Blue",
	}
	electronics := &Electronics{
		BaseProduct: BaseProduct{Name: "Smartphone", Price: 500.0},
		Warranty:    24,
	}

	products := []Product{clothing, electronics}

	fmt.Printf("Total price before discounts: $%.2f\n", CalculateTotalPrice(products))

	clothing.ApplyDiscount(10)   
	electronics.ApplyDiscount(5) 

	fmt.Printf("Total price after discounts: $%.2f\n", CalculateTotalPrice(products))

	clothing.UpdateCharacteristics("L", "Red")
	electronics.UpdateWarranty(36)

	fmt.Printf("Updated Clothing: %+v\n", clothing)
	fmt.Printf("Updated Electronics: %+v\n", electronics)
}
