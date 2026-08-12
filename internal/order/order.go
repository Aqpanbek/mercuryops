package order

import (
	"errors"
)

var ErrInvalidPrice = errors.New("Price need to be positive")
var ErrInvalidQuantity = errors.New("quantity must be positive")
var ErrInvalidID = errors.New("product ID cannot be negative")
var ErrInvalidStatus = errors.New("product status is invalid")

type Order struct {
	id        int64
	sku       int64
	quantity  int64
	unitPrice int64
	status    Status
}

func New(id int64, sku int64, quantity int64, unitPrice int64) (Order, error) {
	if quantity <= 0 {
		return Order{}, ErrInvalidQuantity
	}

	if unitPrice <= 0 {
		return Order{}, ErrInvalidPrice
	}

	return Order{id: id, sku: sku, quantity: quantity, unitPrice: unitPrice, status: StatusDraft}, nil
}

func (o *Order) Confirm() error {

	if o.status != StatusDraft {
		return ErrInvalidStatus
	}

	o.status = StatusConfirmed
	return nil
}

func (o *Order) Cancel() error {
	if o.status != StatusDraft {
		return ErrInvalidStatus
	}

	o.status = StatusCancelled
	return nil
}
