package product

import "testing" // Импортируем тестинг

// создаем функцию для теста.
// t.Fatalf сразу вырубиает тест.
// t.Errorf пишет какая ошибка случилось
func TestNewValidProduct(t *testing.T) {
	item, err := New("Мышь", 15000, 2)

	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}

	if item.Name() != "Мышь" {
		t.Errorf("ожидали название Мышь, получили %s", item.Name())
	}

	if item.Price() != 15000 {
		t.Errorf("Ожидали цену 15000, получили %d", item.Price())
	}

	if item.SKU != 2 {
		t.Errorf("Ожидали ID 2, получили %d", item.SKU)
	}
}

func TestNewRejectsNegativeProducts(t *testing.T) {
	_, err := New("Мышь", -1, 2)

	if err != ErrInvalidPrice {
		t.Errorf("Ожидали ErrInavlidPrice, получили %v", err)
	}

}

func TestNewRejectsEmptyProducts(t *testing.T) {
	_, err := New("", 1500, 2)

	if err != ErrEmptyName {
		t.Errorf("Ожидали ErrEmptyName, получили %v", err)
	}
}

func TestNewRejectsNegativeSkuProducts(t *testing.T) {
	_, err := New("Мышь", 15000, -5)

	if err != ErrInvalidID {
		t.Errorf("Ожидали ErrInvalidID, получили %v", err)
	}
}
