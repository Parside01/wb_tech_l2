package pattern

import "errors"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/**
Область применения паттерна:
	Когда заранее неизвестны тип и зависимости объектов, с которыми должен работать код.

	Когда хотите дать возможность пользователям расширять части вашего кода.

	Когда хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых

Плюсы и минусы этого паттерна:
	Плюсы:
		- Избавляет класс от привязки к конкретным классам продуктов.
		- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
		- Упрощает добавление новых продуктов в программу.
		- Реализует принцип открытости/закрытости.
	Минусы:
		- Если объектов много, то нужно создавать создателей для каждого.
*/

type Gun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) SetName(name string) {
	g.name = name
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) SetPower(power int) {
	g.power = power
}

func (g *gun) GetPower() int {
	return g.power
}

func GetGun(t string) (Gun, error) {
	if t == "Gun" {
		return &gun{
			name:  "Gun",
			power: 10,
		}, nil
	}
	return nil, errors.New("Unknown type")
}
