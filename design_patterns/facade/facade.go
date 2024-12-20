package facade

import (
	"fmt"
	"io"
)

/**
Область применения паттерна:
	Если у нас есть очень сложная подсистема, в данном примере это процессор,
	а нам нужно использовать не полный функционал, а только самые нужные фичи,
	то мы можем использовать фасад, чтобы оставить для использования только нужное,
	а все внутренние сложности оставить за фасадом.

Плюсы и минусы этого паттерна:
	Плюсы:
		- Изоляция клиента от компонентов сложной подсистемы.
	Минусы:
		- При неправильном использовании есть риск превратить
		  такой фасад в "божественный объект", который будет
		  зависеть от всех объектов программы.
*/

type CPU struct {
	writer io.Writer
}

func (c *CPU) Freeze() {
	if _, err := fmt.Fprintln(c.writer, "CPU freeze"); err != nil {
		panic(err)
	}
}

func (c *CPU) Jump(pos int) {
	if _, err := fmt.Fprintln(c.writer, "CPU jump to", pos); err != nil {
		panic(err)
	}
}

func (c *CPU) Execute() {
	if _, err := fmt.Fprintln(c.writer, "CPU execute"); err != nil {
		panic(err)
	}
}

type Memory struct {
	writer io.Writer
}

func (m *Memory) Load(position, data int) {
	if _, err := fmt.Fprintf(m.writer, "Memory load data %d to position %d\n", data, position); err != nil {
		panic(err)
	}
}

type HardDrive struct {
	writer io.Writer
}

func (h *HardDrive) Read(position, size int) {
	if _, err := fmt.Fprintf(h.writer, "Read data from position %d with size %d", position, size); err != nil {
		panic(err)
	}
}

type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputer(w io.Writer) *Computer {
	return &Computer{
		cpu:       &CPU{writer: w},
		memory:    &Memory{writer: w},
		hardDrive: &HardDrive{writer: w},
	}
}

func (c *Computer) Start() {
	c.cpu.Freeze()
	c.memory.Load(0, 2)
	c.cpu.Jump(0)
	c.cpu.Execute()
}
