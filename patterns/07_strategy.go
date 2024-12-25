package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/**
Область применения паттерна:
	Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
	Когда есть множество похожих классов, отличающихся только некоторым поведением.
	Когда не хотим обнажать детали реализации алгоритмов для других классов.

Плюсы:
	- Уход от наследования к делегированию.
	- Реализует принцип открытости/закрытости.
 	- Изолирует код и данные алгоритмов от остальных классов.
Минусы:
	- Усложняет программу за счёт дополнительных классов.
	- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type Strategy interface {
	Print()
}

type FirstStrategy struct{}
type SecondStrategy struct{}

func (FirstStrategy) Print() {
	fmt.Println("First Print")
}

func (SecondStrategy) Print() {
	fmt.Println("Second Print")
}

type Printer struct {
	// Ну то есть в зависимости от условия можем подменять этот объект.
	PrintStrategy Strategy
}

func (p *Printer) Print() {
	p.PrintStrategy.Print()
}