package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
/**
Область применения паттерна:
	Когда у вас есть объект, поведение которого кардинально
	меняется в зависимости от внутреннего состояния, причём типов
	состояний много, и их код часто меняется.

	Когда код объекта содержит множество больших, похожих друг на друга,
	условных операторов, которые выбирают поведения в зависимости от
	текущих значений полей объекта.

Плюсы:
	- Избавляет от множества больших условных операторов.
	- Концентрирует в одном месте код, связанный с определённым состоянием.
Минусы:
	- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type State interface {
	Speak()
}

type Human struct {
	state State
}

func NewHuman(startState State) *Human {
	return &Human{state: startState}
}

func (h *Human) SetState(state State) {
	h.state = state
}

func (h *Human) Speak() {
	h.state.Speak()
}

type SleepState struct{}

func (s *SleepState) Speak() {
	fmt.Println("zzzzzzzz")
}

type WakingState struct{}

func (w *WakingState) Speak() {
	fmt.Println("Hello")
}
