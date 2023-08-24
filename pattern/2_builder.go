package main

import "fmt"

// Интерфейс Строителя объявляет создающие методы для различных частей объектов продуктов
type Builder interface {
	SetPart1(part1 string)
	SetPart2(part2 string)
	SetPart3(part3 string)
}

// Конкретные Строители реализуют методы создания и сборки различных частей продуктов
type ConcreteBuilder struct {
	product *Product
}

func (b *ConcreteBuilder) SetPart1(part1 string) {
	b.product.Part1 = part1
}

func (b *ConcreteBuilder) SetPart2(part2 string) {
	b.product.Part2 = part2
}

func (b *ConcreteBuilder) SetPart3(part3 string) {
	b.product.Part3 = part3
}

// Продукты различных строителей могут не иметь общего интерфейса, поэтому
// можно объединить их в одну структуру, чтобы обращаться к ним через единый интерфейс
type Product struct {
	Part1 string
	Part2 string
	Part3 string
}

// Директор отвечает только за выполнение шагов построения в определённой
// последовательности.
type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.SetPart1("Part 1")
	d.builder.SetPart2("Part 2")
	d.builder.SetPart3("Part 3")
}

func main() {
	builder := &ConcreteBuilder{}
	director := &Director{builder: builder}

	director.Construct()
	product := builder.product

	fmt.Println(product.Part1)
	fmt.Println(product.Part2)
	fmt.Println(product.Part3)
}
