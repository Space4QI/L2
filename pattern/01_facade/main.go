package main

import "fmt"

//Паттерн "Фасад" (Facade) предоставляет унифицированный интерфейс к набору интерфейсов в подсистеме.
//Фасад определяет высокоуровневый интерфейс, который упрощает использование подсистемы.
//
//Паттерн "Фасад" используется, когда:
//
//- Нужно предоставить простой интерфейс к сложной системе.
//- Необходимо декомпозировать подсистемы в независимые части.
//- Требуется уменьшить количество зависимостей между клиентом и сложной системой.
//- Существует необходимость предоставить упрощенный интерфейс к библиотеке или фреймворку.
//Плюсы:
//
//- Упрощение использования: обеспечивает простой интерфейс к сложной подсистеме.
//- Слабое связывание: уменьшает количество зависимостей между клиентом и сложной системой.
//- Инкапсуляция деталей: скрывает детали реализации подсистемы от клиентов.
//- Гибкость: можно изменить подсистему без изменения интерфейса фасада.
//Минусы:
//
//- Дополнительный слой: вводит дополнительный уровень абстракции, который может немного замедлить систему.
//- Не подходит для всех задач: если интерфейс слишком упрощен, могут быть потеряны возможности полной
//функциональности подсистемы.
//Примеры использования
//GUI библиотеки: в графических интерфейсах, таких как Qt или Swing, фасад используется для упрощения работы
//с окнами, кнопками и другими компонентами.
//Системы сборки: в инструментах сборки, таких как Maven или Gradle, фасад обеспечивает простой интерфейс
//для конфигурации и запуска сборок, скрывая сложные детали конфигурации.
//Web Frameworks: в веб-фреймворках, таких как Django или Spring, фасад может использоваться
//для упрощения работы с маршрутизацией, контроллерами и базой данных.
//Медиаплееры: в медиаплеерах фасад может скрывать сложные операции работы с различными форматами
//медиафайлов и предоставлять простой интерфейс для воспроизведения, паузы и остановки.
//Примером использования паттерна "Фасад" может быть разработка системы домашнего кинотеатра.
//Подсистема домашнего кинотеатра включает в себя DVD-плеер, телевизор, звуковую систему и освещение.
//Фасад создаст простой интерфейс для управления всей системой.

// Компонент A
type ComponentA struct{}

func (c *ComponentA) OperationA1() {
	fmt.Println("Component A, Method OperationA1")
}

func (c *ComponentA) OperationA2() {
	fmt.Println("Component A, Method OperationA2")
}

// Компонент B
type ComponentB struct{}

func (c *ComponentB) OperationB1() {
	fmt.Println("Component B, Method OperationB1")
}

func (c *ComponentB) OperationB2() {
	fmt.Println("Component B, Method OperationB2")
}

// Компонент C
type ComponentC struct{}

func (c *ComponentC) OperationC1() {
	fmt.Println("Component C, Method OperationC1")
}

func (c *ComponentC) OperationC2() {
	fmt.Println("Component C, Method OperationC2")
}

// Фасад
type Facade struct {
	componentA *ComponentA
	componentB *ComponentB
	componentC *ComponentC
}

func NewFacade() *Facade {
	return &Facade{
		componentA: &ComponentA{},
		componentB: &ComponentB{},
		componentC: &ComponentC{},
	}
}

func (f *Facade) Operation1() {
	fmt.Println("Facade Operation1:")
	f.componentA.OperationA1()
	f.componentB.OperationB1()
}

func (f *Facade) Operation2() {
	fmt.Println("Facade Operation2:")
	f.componentA.OperationA2()
	f.componentB.OperationB2()
	f.componentC.OperationC1()
}

func (f *Facade) Operation3() {
	fmt.Println("Facade Operation3:")
	f.componentC.OperationC2()
}

func main() {
	facade := NewFacade()

	facade.Operation1()
	fmt.Println()
	facade.Operation2()
	fmt.Println()
	facade.Operation3()
}
