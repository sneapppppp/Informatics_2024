package lab7

import (
	"fmt"
)

func Executelab7() {
	laptop := &laptop{Name: "Laptop", Price: 1000}
	phone := &phone{Name: "Phone", Price: 500}
	tablet := &tablet{Name: "Tablet", Price: 300}

	products := []Product{laptop, phone, tablet}

	fmt.Printf("Total before discount: %.2f\n", CalculateTotal(products))

	for _, product := range products {
		product.ApplyDiscount(10)
	}

	fmt.Printf("Total after discount: %.2f\n", CalculateTotal(products))
}
