package sixthhw

import "fmt"

// 1

type Book struct {
	Title       string
	Author      string
	pages       int
	isAvailable bool
}

func newBook(title string, author string, pages int) *Book {
	return &Book{
		Title:       title,
		Author:      author,
		pages:       pages,
		isAvailable: true,
	}
}

func (b *Book) Info() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Pages:", b.pages)
	fmt.Println("Available:", b.isAvailable)
}

func (b *Book) Borrow() {
	if b.isAvailable {
		b.isAvailable = false
		fmt.Println("Книга выдана")
	} else {
		fmt.Println("Книга недоступна")
	}
}

func (b *Book) ReturnBook() {
	b.isAvailable = true
	fmt.Println("Книга возвращена")
}

func (b *Book) GetPages() int {
	return b.pages
}

func (b *Book) SetPages(p int) {
	if p <= 0 {
		fmt.Println("Ошибка: количество страниц должно быть > 0")
		return
	}
	b.pages = p
}

// 2

type Worker interface {
	Work() string
	GetName() string
}

type Programmer struct {
	Name     string
	Language string
}

func (p Programmer) Work() string {
	return fmt.Sprintf("Программист %s пишет код на %s", p.Name, p.Language)
}

func (p Programmer) GetName() string {
	return p.Name
}

type Designer struct {
	Name string
	Tool string
}

func (d Designer) Work() string {
	return fmt.Sprintf("Дизайнер %s делает макет в %s", d.Name, d.Tool)
}

func (d Designer) GetName() string {
	return d.Name
}

func ShowWork(w Worker) {
	fmt.Println("Имя:", w.GetName())
	fmt.Println(w.Work())
}

// 3

type Product struct {
	Name     string
	price    float64
	Quantity int
}

func newProduct(name string, price float64, quantity int) Product {
	if price < 0 {
		price = 0
	}
	if quantity < 0 {
		quantity = 0
	}
	return Product{
		Name:     name,
		price:    price,
		Quantity: quantity,
	}
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetPrice(newPrice float64) {
	if newPrice < 0 {
		fmt.Println("Ошибка: цена не может быть отрицательной")
		return
	}
	p.price = newPrice
}

func (p *Product) Buy(amount int) {
	if amount <= 0 {
		fmt.Println("Ошибка: некорректное количество")
		return
	}
	if p.Quantity < amount {
		fmt.Println("Недостаточно товара")
		return
	}
	p.Quantity -= amount
	fmt.Println("Покупка успешна")
}

func (p *Product) Restock(amount int) {
	if amount <= 0 {
		fmt.Println("Ошибка: некорректное количество")
		return
	}
	p.Quantity += amount
}

func (p Product) Info() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Price:", p.price)
	fmt.Println("Quantity:", p.Quantity)
}

func main() {
}