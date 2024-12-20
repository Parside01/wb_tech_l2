package command

/**
Область применения паттерна:
	Когда есть несколько объектов, которые могут выполнять одну и ту же логику.

	Мы хотим ставить операции в очередь, выполнять по
	расписанию или передавать по сети.

	Если, вдруг, понадобится отмена операций.

Плюсы и минусы этого паттерна:
	Плюсы:
		- Убирает зависимость между объектами, вызывающими операции, и объектами, которые их выполняют.
		- Простая реализация отмены и повтора операций.
		- Отложенный запуск операции.
		- Собирать сложные команды из простых.
		- Реализует принцип открытости для расширения\закрытости для изменения.
	Минусы:
		- Усложняет код из-за порождения новых сущностей.
*/

type Command interface {
	Execute()
}

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

type Device interface {
	On()
	Off()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}

type TV struct {
	Running bool
}

func (t *TV) On() {
	t.Running = true
}

func (t *TV) Off() {
	t.Running = false
}
