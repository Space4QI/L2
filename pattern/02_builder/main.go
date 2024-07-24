package main

import "fmt"

//Паттерн "Строитель" (Builder) предназначен для пошагового конструирования сложных объектов.
//Он позволяет создавать различные представления одного и того же объекта, отделяя процесс
//конструирования от представления.
//
//Применимость
//Когда объект имеет много параметров: если конструктор объекта содержит слишком много параметров,
//паттерн "Строитель" помогает избежать проблем с читаемостью и поддерживаемостью кода.
//Когда нужно создавать разные представления одного и того же объекта: паттерн позволяет создавать
//различные представления одного и того же объекта с помощью одного и того же процесса конструирования.
//Когда необходимо обеспечить пошаговое создание объекта: паттерн позволяет пошагово
//создавать объект, что облегчает его модификацию на разных этапах.
//Плюсы
//Разделение сложного конструктора: упрощает создание сложных объектов, разделяя процесс на отдельные шаги.
//Изоляция кода: строитель изолирует код для создания и представления объекта.
//Гибкость: позволяет создавать разные представления одного и того же объекта, используя один и
//тот же процесс конструирования.
//Минусы
//Дополнительная сложность: добавляет дополнительные классы и методы, что может усложнить систему.
//Трудоемкость:реализация паттерна может потребовать написания значительного количества кода.
//Реальные примеры использования
//Библиотеки GUI: паттерн "Строитель" часто используется для создания сложных графических интерфейсов,
//где необходимо пошагово добавлять элементы.
//Создание отчетов: в системах создания отчетов паттерн используется для пошагового добавления различных
//частей отчета (заголовки, таблицы, графики и т.д.).
//Конфигурационные системы: паттерн применяется для создания сложных конфигурационных объектов,
//которые могут иметь множество параметров.

// House - продукт, который будет создаваться строителем
type House struct {
	Windows string
	Doors   string
	Roof    string
	Garage  string
	Pool    string
}

func (h *House) Show() {
	fmt.Printf("House with %s windows, %s doors, %s roof, %s garage, %s pool\n", h.Windows, h.Doors, h.Roof, h.Garage, h.Pool)
}

// HouseBuilder - интерфейс строителя
type HouseBuilder interface {
	SetWindows() HouseBuilder
	SetDoors() HouseBuilder
	SetRoof() HouseBuilder
	SetGarage() HouseBuilder
	SetPool() HouseBuilder
	Build() *House
}

// ConcreteHouseBuilder - конкретный строитель для дома
type ConcreteHouseBuilder struct {
	house *House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{house: &House{}}
}

func (b *ConcreteHouseBuilder) SetWindows() HouseBuilder {
	b.house.Windows = "Double-glazed"
	return b
}

func (b *ConcreteHouseBuilder) SetDoors() HouseBuilder {
	b.house.Doors = "Wooden"
	return b
}

func (b *ConcreteHouseBuilder) SetRoof() HouseBuilder {
	b.house.Roof = "Shingle"
	return b
}

func (b *ConcreteHouseBuilder) SetGarage() HouseBuilder {
	b.house.Garage = "Two-car"
	return b
}

func (b *ConcreteHouseBuilder) SetPool() HouseBuilder {
	b.house.Pool = "Outdoor"
	return b
}

func (b *ConcreteHouseBuilder) Build() *House {
	return b.house
}

// Director - директор, который управляет процессом строительства
type Director struct {
	builder HouseBuilder
}

func NewDirector(builder HouseBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() *House {
	return d.builder.SetWindows().
		SetDoors().
		SetRoof().
		SetGarage().
		SetPool().
		Build()
}

func main() {
	builder := NewConcreteHouseBuilder()
	director := NewDirector(builder)
	house := director.Construct()
	house.Show()
}
