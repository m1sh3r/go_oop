# РЕАЛИЗАЦИЯ ООП В GO

| Концепция    | В Go                                                           |
| ------------ | -------------------------------------------------------------- |
| Абстракция   | Структуры (struct)                                             |
| Инкапсуляция | Экспортируемые/неэкспортируемые поля (нэйминг с большой буквы) |
| Наследование | Композиция (встраивание)                                       |
| Полиморфизм  | Интерфейсы (interface)                                         |

# СТРУКТУРЫ

Структуры — основной способ описания данных.

[Пример](1_struct/main.go)

```go
type Animal struct {
Name    string
Species string
Age     int
}
```

# МЕТОДЫ

Методы - функции, привязанные к структуре

[Пример](2_method/person/person.go)

```go
type Person struct {
FirstName string
LastName  string
}

func (p Person) Introduce() string {
return fmt.Sprintf("Привет, я %s %s", p.FirstName, p.LastName)
}

```

Разница значения vs указателя:

- Методы со значением → копия структуры.
- Методы с указателем → изменяют исходный объект.

```go
type Person struct {
FirstName string
LastName  string
}

func (p Person) Introduce() string {
return fmt.Sprintf("Привет, я %s %s", p.FirstName, p.LastName)
}

```

# ИНКАПСУЛЯЦИЯ

В Go нет ключевых слов private/public.
Доступ определяется именованием:

- с заглавной буквы → экспортируемое (public).

- с маленькой буквы → приватное (private).

[Пример](3_encapsulation/person/person.go)

```go
package person1

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) GetFirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(name string) {
	p.firstName = name
}

func (p *Person) SetLastName(name string) {
	p.lastName = name
}

func (p *Person) GetLastName() string {
	return p.lastName
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("Привет, я %s %s", p.firstName, p.lastName)
}

```

# КОМПОЗИЦИЯ

Go не поддерживает наследование, но есть композиция (встраивание).

[Пример](4_%D1%81omposition/person/person.go)

```go
type Address struct {
City string
}

type Person struct {
Name string
Address
}

```

# ИНТЕРФЕЙСЫ И ПОЛИМОРФИЗМ

Интерфейсы позволяют описывать поведение.

````go
type Shape interface {
Area() float64
}

type Circle struct {
Radius float64
}

func (c Circle) Area() float64 {
return 3.14 * c.Radius * c.Radius
}

# Выравнивание

Плохое выравнивание

```go
type Bad struct {
	A byte   // 1 байт
	B int64  // 8 байт
	C byte   // 1 байт
	D int32  // 4 байта
}
````

Как это работает:

- A занимает 1 байт → Go выравнивает следующее поле (int64) по 8 байтам → добавляется 7 байт паддинга.

- B = 8 байт.

- C = 1 байт → но дальше идёт int32, который должен быть кратен 4 → добавляется 3 байта паддинга.

- D = 4 байта.
  Итог: 24 байта, хотя данных всего 14 байт.

**Хорошее выравнивание**

```go
type Good struct {
	B int64  // 8 байт
	D int32  // 4 байта
	A byte   // 1 байт
	C byte   // 1 байт
}
```

Как это работает:

- B = 8 байт.

- D = 4 байта.

- A = 1 байт.

- C = 1 байт → всего 14 байт данных.

- Чтобы всё выровнялось на 8 байт, Go добавит только 2 байта паддинга.

Итог: 16 байт.
[Пример](6_alignment/main.go)
