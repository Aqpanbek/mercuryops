package inventory

import "errors"

var ErrInvalidQuantity = errors.New("quantity must be positive")
var ErrInsufficientStock = errors.New("reserved not correct")
var ErrNotEnoughReserved = errors.New("Reserved is not enough")

type Stock struct {
	sku      int64
	onHand   int64
	reserved int64
} // обьявляем структуру Сток с SKU, остатками и резервом

func (s *Stock) Receive(quantity int64) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	s.onHand += quantity
	return nil
}

// создаем функцию которая меняет значение onHand.

func (s Stock) Available() int64 {
	return s.onHand - s.reserved
}

func (s *Stock) Reserve(quantity int64) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	if quantity > s.Available() {
		return ErrInsufficientStock
	}

	s.reserved += quantity
	return nil
}

func (s *Stock) Release(quantity int64) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	if quantity > s.reserved {
		return ErrNotEnoughReserved
	}

	s.reserved -= quantity
	return nil
}

func (s *Stock) Ship(quantity int64) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}

	if quantity > s.reserved {
		return ErrNotEnoughReserved
	}

	s.onHand -= quantity
	s.reserved -= quantity
	return nil
}
