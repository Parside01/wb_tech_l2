package builder

/**
Область применения паттерна:
	У нас есть сложный объект, который требует сложной пошаговой инициализации
	множества полей и вложенных объектов. Чтобы не распылять этот код в разные
	области программы лучше поручить инициализацию этого объекта строителям.

	Например:
		- Избавиться от "телескопического" конструктора
          NewPizzaDefault(), NewPizzaWithCheese(size int, cheese bool), NewPizzaWithPeperoniAndCheese(size int, cheese bool, pepperoni bool).
		- Создание сложные составные объекты
Плюсы и минусы этого паттерна:
	Плюсы:
		- Создание объектов пошагово.
		- Использование одного и того же кода для создания разных объектов.
        - Изоляция сложной логики создания объекта.
    Минусы:
		- Усложнение кода из-за порождения новых сущностей.
        - Привязка пользователя к конкретным строителям.
*/

type Pizza struct {
	dough    string
	sauce    string
	cheese   string
	toppings []string
}

type PizzaBuilder interface {
	SetDough(dough string)
	SetSauce(sauce string)
	SetCheese(cheese string)
	SetToppings(toppings ...string)
	Build() *Pizza
}

type DefaultPizzaBuilder struct {
	pizza *Pizza
}

func NewDefaultPizzaBuilder() PizzaBuilder {
	return &DefaultPizzaBuilder{pizza: &Pizza{}}
}

func (p *DefaultPizzaBuilder) SetDough(dough string) {
	p.pizza.dough = dough
}

func (p *DefaultPizzaBuilder) SetSauce(sauce string) {
	p.pizza.sauce = sauce
}

func (p *DefaultPizzaBuilder) SetCheese(cheese string) {
	p.pizza.cheese = cheese
}

func (p *DefaultPizzaBuilder) SetToppings(toppings ...string) {
	p.pizza.toppings = toppings
}

func (p *DefaultPizzaBuilder) Build() *Pizza {
	return p.pizza
}

type Director struct {
	builder PizzaBuilder
}

func NewDirector(builder PizzaBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *Pizza {
	d.builder.SetDough("Thin")
	d.builder.SetSauce("Tomato")
	d.builder.SetCheese("Mozzarella")
	d.builder.SetToppings("Mushrooms", "Olives", "Onions")
	return d.builder.Build()
}
