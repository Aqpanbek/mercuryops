package product // При создании файла на GO в начале всегда указать писать какому пакету он принадлежит. К примеру product.go принадлежит пакету product

import "errors" // импортируем пакет ошибок

var ErrEmptyName = errors.New("product name cannot be empty")        // Создаем вид ошибки если нет название у продукта
var ErrInvalidPrice = errors.New("product price cannot be negative") // Так же и с ценой
var ErrInvalidID = errors.New("product ID cannot be negative")       // Так же и с SKU(это ID)

type Product struct {
	name  string
	price int64
	SKU   int64
} // Указываем какие значения имеет продукт

func New(name string, price int64, sku int64) (Product, error) {
	if name == "" {
		return Product{}, ErrEmptyName
	}

	if price < 0 {
		return Product{}, ErrInvalidPrice
	}

	if sku < 0 {
		return Product{}, ErrInvalidID
	}

	return Product{name: name, price: price, SKU: sku}, nil
} // Так как у нас нет возможности создаваиь напрямую, мы пишем функцию которое создает продукт

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() int64 {
	return p.price
}

func (p Product) Sku() int64 {
	return p.SKU
}
