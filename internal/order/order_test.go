package order

import "testing"

func TestNewCreateDraftOrder(t *testing.T) {
	got, err := New(101, 2, 3, 1500)

	if err != nil {
		t.Fatalf("не ожидали ошибку %v", err)
	}

	if got.status != StatusDraft {
		t.Errorf("ожидали статус %s, получили %s", StatusDraft, got.status)
	}

	if got.id != 101 {
		t.Errorf("ожидали id %d, получили %d", 101, got.id)
	}

	if got.sku != 2 {
		t.Errorf("ожидали sku %d, получили %d", 2, got.sku)
	}

	if got.quantity != 3 {
		t.Errorf("ожидали quantity %d, получили %d", 3, got.quantity)
	}

	if got.unitPrice != 1500 {
		t.Errorf("ожидали unitPrice %d, получили %d", 1500, got.unitPrice)
	}

}

func TestNewRejectsZeroQuantity(t *testing.T) {
	_, err := New(101, 2, 0, 1500)

	if err != ErrInvalidQuantity {
		t.Errorf("ожидали ErrInvalidQuantity, получили  %v", err)
	}
}

func TestNewRejectsNegativeQuantity(t *testing.T) {
	_, err := New(101, 2, -1, 1500)

	if err != ErrInvalidQuantity {
		t.Errorf("Ожидали получить ErrInvalidQuantity, получили %v", err)
	}
}

func TestNewRejectsZeroUnitPrice(t *testing.T) {
	_, err := New(101, 2, 1, 0)

	if err != ErrInvalidPrice {
		t.Errorf("ожидали получить ErrInvalidPrice, получили %v", err)
	}
}

func TestNewRejectsNegativeUnitPrice(t *testing.T) {
	_, err := New(101, 2, 1, -1500)

	if err != ErrInvalidPrice {
		t.Errorf("Ожидали получить ErrInvalidPrice, получили %v", err)
	}
}

func TestConfirmDraftOrder(t *testing.T) {
	createdOrder, err := New(101, 2, 1, 1500)

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
	createdOrder, err := New(101, 2, 1, 1500)

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
	createdOrder, err := New(101, 2, 1, 1500)

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
