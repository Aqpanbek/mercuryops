package inventory

import "testing"

func TestReceiveAddsStock(t *testing.T) {
	stock := Stock{sku: 2}

	err := stock.Receive(10)
	if err != nil {
		t.Fatalf("не ождилаи ошибку, %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("ожидали onHand 10, получили %d", stock.onHand)
	}
}

func TestRecieveRejectsNegative(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10}

	err := stock.Receive(-5)

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали ErrInvalidQuantity, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("onHand изменился, ожидали 10, получили %d", stock.onHand)
	}
}

func TestAvailableStock(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 3}

	available := stock.Available()

	if available != 7 {
		t.Errorf("ожидали available 7, получили %d", available)
	}

}

func TestReservRejectsToolargeQuantity(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 3}

	err := stock.Reserve(8)

	if err != ErrInsufficientStock {
		t.Errorf("ожидали ErrInsufficientStock, получили %v", err)
	}

	if stock.reserved != 3 {
		t.Errorf("ожидали reserved 3, получили %d", stock.reserved)
	}
}

func TestReceiveRejectsZero(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10}

	err := stock.Receive(0)

	if err != ErrInvalidQuantity {
		t.Errorf("ожидали ErrInvalidQuantity, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("onHand изменился: ожидали 10, получили %d", stock.onHand)
	}
}

func TestReserveRejectsZero(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 3}

	err := stock.Reserve(0)

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали ErrInvalidQuantity, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("onHand изменился: ожидали 8, получили %d", stock.onHand)
	}

	if stock.reserved != 3 {
		t.Errorf("reserved изменился: ожидали 3, получили %d", stock.reserved)
	}
}

func TestReserveStock(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 0}

	err := stock.Reserve(4)

	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}

	if stock.reserved != 4 {
		t.Errorf("ожидали reserved 4, получили %d", stock.reserved)
	}

	if stock.onHand != 10 {
		t.Errorf("ожидали onHand 10, получили %d", stock.onHand)
	}

	available := stock.Available()

	if available != 6 {
		t.Errorf("ожидали available 6, получили %d", available)
	}
}

func TestRelease(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 4}

	err := stock.Release(2)

	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}

	if stock.reserved != 2 {
		t.Errorf("ожидали reserved 2, получили %d", stock.reserved)
	}

	if stock.onHand != 10 {
		t.Errorf("ожидали onHand 10, получили %d", stock.onHand)
	}

	available := stock.Available()

	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}
}

func TestShip(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 4}

	err := stock.Ship(2)

	if err != nil {
		t.Fatalf("не ожидали ошибку: %v", err)
	}

	if stock.reserved != 2 {
		t.Errorf("Ожидали reserved 2, получили %d", stock.reserved)
	}

	if stock.onHand != 8 {
		t.Errorf("Ожидали onHand 8, получили %d", stock.onHand)
	}

	available := stock.Available()

	if available != 6 {
		t.Errorf("Ожидали available 6, получили %d", available)
	}
}

func TestShipRejectsQuantityAboveReserved(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 2}

	err := stock.Ship(4)

	if err != ErrNotEnoughReserved {
		t.Errorf("ожидали ErrNotEnoughReserved, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 2 {
		t.Errorf("Ожидали reserved 2, получили %d", stock.reserved)
	}

	available := stock.Available()
	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}
}

func TestShipRejectsNegativeQuantity(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 2}

	err := stock.Ship(-1)

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали ErrInvalidQuantity, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 2 {
		t.Errorf("Ожидали reserved 2, получили %d", stock.reserved)
	}

	available := stock.Available()
	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}
}

func TestShipRejectsZeroQuantity(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 2}

	err := stock.Ship(0)

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали ErrInvalidQuantity, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 2 {
		t.Errorf("Ожидали reserved 2, получили %d", stock.reserved)
	}

	available := stock.Available()
	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}
}

func TestReleaseRejectsQuantityAboveReserved(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 2}

	err := stock.Release(4)

	if err != ErrNotEnoughReserved {
		t.Errorf("Ожидали ErrNotEnoughReserved, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 2 {
		t.Errorf("Ожидали reserved 2, получили %d", stock.reserved)
	}

	available := stock.Available()
	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}
}

func TestReleaseRejectsQuantityZero(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 2}

	err := stock.Release(0)

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали ErrInvalidQuantity, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 2 {
		t.Errorf("Ожидали reserved 2, получили %d", stock.reserved)
	}

	available := stock.Available()
	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}
}

func TestReleaseRejectsQuantityNegative(t *testing.T) {
	stock := Stock{sku: 2, onHand: 10, reserved: 2}

	err := stock.Release(-1)

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали ErrInvalidQuantity, получили %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 2 {
		t.Errorf("Ожидали reserved 2, получили %d", stock.reserved)
	}

	available := stock.Available()
	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}
}

func TestStockLifecycle(t *testing.T) {
	stock := Stock{sku: 2, onHand: 0, reserved: 0}

	err := stock.Receive(10)
	if err != nil {
		t.Fatalf("не ождилаи ошибку, %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("ожидали onHand 10, получили %d", stock.onHand)
	}

	err = stock.Reserve(3)

	if err != nil {
		t.Fatalf("Не ожидали ошибку, %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 3 {
		t.Errorf("Ожидали reserved 3, получили %d", stock.reserved)
	}

	available := stock.Available()

	if available != 7 {
		t.Errorf("Ожидали available 7, получили %d", available)
	}

	err = stock.Release(1)
	if err != nil {
		t.Fatalf("не ожидали ошибку, %v", err)
	}

	if stock.onHand != 10 {
		t.Errorf("Ожидали onHand 10, получили %d", stock.onHand)
	}

	if stock.reserved != 2 {
		t.Errorf("ожидали reserved 2, получили %d", stock.reserved)
	}

	available = stock.Available()

	if available != 8 {
		t.Errorf("ожидали available 8, получили %d", available)
	}

	err = stock.Ship(2)

	if err != nil {
		t.Fatalf("Не ожидали ошибку %v", err)
	}

	if stock.onHand != 8 {
		t.Errorf("Ожидали onHand 8, получили %d", stock.onHand)
	}

	if stock.reserved != 0 {
		t.Errorf("Ожидали reserved 0, получили %d", stock.reserved)
	}

	available = stock.Available()

	if available != 8 {
		t.Errorf("Ожидали available 8, получили %d", available)
	}

}
