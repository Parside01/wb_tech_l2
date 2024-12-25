package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
/**
Область применения паттерна:
	Когда программа должна обрабатывать разные запросы несколькими способами,
	но заранее не известно какие запросы будут приходить и какие обработчики для них потребуются.

	Когда важно, чтобы обработчики выполнялись в строгом порядке.

	Когда набор объектов, способных обработать запрос должен задаваться динамически.

Плюсы и минусы этого паттерна:
	Плюсы:
 		- Уменьшает зависимость между клиентами и обработчиками.
		- Реализует принцип единой ответственности.
		- Реализует принцип открытости для расширения\закрытости для изменения.
	Минусы:
		- Запрос может остаться никем не обработанным.
*/

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Handler interface {
	Execute(*Patient)
	SetNext(next Handler)
}

type Reception struct {
	next Handler
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	} else if r.next != nil {
		fmt.Println("Reception registration patient")
		p.registrationDone = true
		r.next.Execute(p)
	}

}

func (r *Reception) SetNext(next Handler) {
	r.next = next
}

type Doctor struct {
	next Handler
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Patient doctor already checked")
		d.next.Execute(p)
		return
	} else if p.doctorCheckUpDone {
		fmt.Println("Doctor doctor patient")
		d.next.Execute(p)
		p.doctorCheckUpDone = true
	}
}

func (d *Doctor) SetNext(next Handler) {
	d.next = next
}

type Medical struct {
	next Handler
}

func (m *Medical) Execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medical patient already done")
		m.next.Execute(p)
		return
	} else if m.next != nil {
		fmt.Println("Medical medical patient")
		m.next.Execute(p)
		p.medicineDone = true
	}

}

func (m *Medical) SetNext(next Handler) {
	m.next = next
}

type Cashier struct {
	next Handler
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Patient payment already done")
		c.next.Execute(p)
		return
	} else if c.next != nil {
		fmt.Println("Patient payment patient")
		c.next.Execute(p)
		p.paymentDone = true
	}
}
