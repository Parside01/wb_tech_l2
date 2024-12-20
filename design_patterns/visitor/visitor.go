package visitor

import (
	"math"
)

/**
Область применения паттерна:
	Если нам нужно выполнить одинаковую операцию
	над объектом сложной структуры, при этом эти
	объекты принадлежат к разным типам.

	Если мы хотим выполнить над объектами операции, но
	эти операции не связаны с самим объектом.

	Применить новое поведение для конкретных классов их текущей
	иерархии, а затрагивая другие.

Плюсы и минусы этого паттерна:
	Плюсы:
		- Упрощает добавление новых операций, работающих со
          текущими объектами.
		- Объединяет родственные операции в одном объекте.
		- Посетитель может накапливать состояние при обходе структуры объектов.
	Минусы:
		- Нарушение инкапсуляции.
*/

type Visitor interface {
	VisitForCircle(*Circle)
	VisitForSquare(*Square)
	VisitForRectangle(*Rectangle)
}

type Shape interface {
	accept(visitor Visitor)
}

type Circle struct {
	Radius float64
}

func (c *Circle) accept(v Visitor) {
	v.VisitForCircle(c)
}

type Square struct {
	Side float64
}

func (s *Square) accept(v Visitor) {
	v.VisitForSquare(s)
}

type Rectangle struct {
	Left, Right float64
}

func (r *Rectangle) accept(v Visitor) {
	v.VisitForRectangle(r)
}

type AreaCalc struct {
	Area float64
}

func (a *AreaCalc) VisitForCircle(c *Circle) {
	a.Area = c.Radius * c.Radius * math.Pi
}

func (a *AreaCalc) VisitForSquare(s *Square) {
	a.Area = s.Side * s.Side
}

func (a *AreaCalc) VisitForRectangle(r *Rectangle) {
	a.Area = r.Right * r.Left
}
