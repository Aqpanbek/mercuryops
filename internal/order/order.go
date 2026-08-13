package order

import (
	"errors"
)

var ErrInvalidPrice = errors.New("Price need to be positive")
var ErrInvalidQuantity = errors.New("quantity must be positive")
var ErrInvalidID = errors.New("product ID cannot be negative")
var ErrInvalidStatus = errors.New("product status is invalid")

type Order struct {
	id     int64
	items  []OrderItem
	status Status
}

func New(id int64, items []OrderItem) (Order, error) {
	for _, item := range items {
		if item.quantity <= 0 {
			return Order{}, ErrInvalidQuantity
		}

		if item.unitPrice <= 0 {
			return Order{}, ErrInvalidPrice
		}
	}

	return Order{
		id:     id,
		items:  items,
		status: StatusDraft,
	}, nil
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
