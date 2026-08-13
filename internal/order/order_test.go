package order

import "testing"

func TestNewCreateDraftOrder(t *testing.T) {
	items := []OrderItem{
		{sku: 2, quantity: 3, unitPrice: 1500},
		{sku: 3, quantity: 1, unitPrice: 500},
	}

	got, err := New(101, items)

	if err != nil {
		t.Fatalf("не ожидали ошибку %v", err)
	}

	if got.status != StatusDraft {
		t.Errorf("ожидали статус %s, получили %s", StatusDraft, got.status)
	}

	if got.id != 101 {
		t.Errorf("ожидали id %d, получили %d", 101, got.id)
	}

	if len(got.items) != 2 {
		t.Fatalf("ожидали 2 позиции, получили %d", len(got.items))
	}

	if got.items[0].sku != 2 {
		t.Errorf("ожидали sku %d, получили %d", 2, got.items[0].sku)
	}

	if got.items[0].quantity != 3 {
		t.Errorf("ожидали quantity %d, получили %d", 3, got.items[0].quantity)
	}

	if got.items[0].unitPrice != 1500 {
		t.Errorf("ожидали unitPrice %d, получили %d", 1500, got.items[0].unitPrice)
	}
}

func TestNewRejectsZeroQuantity(t *testing.T) {
	_, err := New(101, []OrderItem{
		{sku: 2, quantity: 1, unitPrice: 1500},
		{sku: 3, quantity: 0, unitPrice: 500},
	})

	if err != ErrInvalidQuantity {
		t.Errorf("ожидали ErrInvalidQuantity, получили  %v", err)
	}
}

func TestNewRejectsNegativeQuantity(t *testing.T) {
	_, err := New(101, []OrderItem{{sku: 2, quantity: -1, unitPrice: 1500}})

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали получить ErrInvalidQuantity, получили %v", err)
	}
}

func TestNewRejectsZeroUnitPrice(t *testing.T) {
	_, err := New(101, []OrderItem{{sku: 2, quantity: 1, unitPrice: 0}})

	if err != ErrInvalidPrice {
		t.Errorf("ожидали получить ErrInvalidPrice, получили %v", err)
	}
}

func TestNewRejectsNegativeUnitPrice(t *testing.T) {
	_, err := New(101, []OrderItem{{sku: 2, quantity: 1, unitPrice: -1500}})

	if err != ErrInvalidPrice {
		t.Errorf("Ожидали получить ErrInvalidPrice, получили %v", err)
	}
}

func TestConfirmDraftOrder(t *testing.T) {
	createdOrder, err := New(101, []OrderItem{{sku: 2, quantity: 1, unitPrice: 1500}})

	if err != nil {
		t.Fatalf("не ожадила ошибку, получили %v", err)
	}

	err = createdOrder.Confirm()

	if err != nil {
		t.Fatalf("не ожадила ошибку, получили %v", err)
	}

	if createdOrder.status != StatusConfirmed {
		t.Errorf("Ожидали StatusConfirmed, получили %v", createdOrder.status)
	}
}

func TestConfirmRejectsAlreadyConfirmedOrder(t *testing.T) {
	createdOrder, err := New(101, []OrderItem{{sku: 2, quantity: 1, unitPrice: 1500}})

	if err != nil {
		t.Fatalf("не ожадила ошибку, получили %v", err)
	}

	err = createdOrder.Confirm()

	if err != nil {
		t.Fatalf("не ожадила ошибку, получили %v", err)
	}

	err = createdOrder.Confirm()

	if err != ErrInvalidStatus {
		t.Errorf("ожидали, ErrInvalidStatus, получили %v", err)
	}

	if createdOrder.status != StatusConfirmed {
		t.Errorf("ожидали статус %s, получили %s",
			StatusConfirmed,
			createdOrder.status,
		)
	}
}

func TestConfirmRejectsCancelledOrder(t *testing.T) {
	createdOrder, err := New(101, []OrderItem{{sku: 2, quantity: 1, unitPrice: 1500}})

	if err != nil {
		t.Fatalf("не ожидали ошибку, получили %v", err)
	}

	createdOrder.status = StatusCancelled

	err = createdOrder.Confirm()

	if err != ErrInvalidStatus {
		t.Errorf("Ожидали ErrInvalidStatus, получили %v", err)
	}

	if createdOrder.status != StatusCancelled {
		t.Errorf("Ожидали статус %s, получили %s", StatusCancelled, createdOrder.status)
	}

}
