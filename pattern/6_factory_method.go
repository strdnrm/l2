package main

import "fmt"

// Интерфейс продукта
type Product interface {
	GetName() string
}

type ProductA struct{}

func (p *ProductA) GetName() string {
	return "Product A"
}

type ProductB struct{}

func (p *ProductB) GetName() string {
	return "Product B"
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

type FactoryA struct{}

func (f *FactoryA) CreateProduct() Product {
	return &ProductA{}
}

type FactoryB struct{}

func (f *FactoryB) CreateProduct() Product {
	return &ProductB{}
}

func main() {
	factoryA := &FactoryA{}
	productA := factoryA.CreateProduct()
	fmt.Println(productA.GetName())

	factoryB := &FactoryB{}
	productB := factoryB.CreateProduct()
	fmt.Println(productB.GetName())
}
