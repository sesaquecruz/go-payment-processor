package entity

import (
	"errors"

	app_errors "github.com/sesaquecruz/go-payment-processor/internal/application/errors"
)

type Transaction struct {
	Card     Card     `json:"card"`
	Purchase Purchase `json:"purchase"`
	Store    Store    `json:"store"`
	Acquirer Acquirer `json:"-"`
}

func NewTransaction(card Card, purchase Purchase, store Store, acquirer Acquirer) *Transaction {
	return &Transaction{
		Card:     card,
		Purchase: purchase,
		Store:    store,
		Acquirer: acquirer,
	}
}

func (t *Transaction) Validate() error {
	errs := make([]error, 0)

	if err := t.Card.Validate(); err != nil {
		var v *app_errors.Validation
		if errors.As(err, &v) {
			errs = append(errs, v.Unwrap()...)
		}
	}

	if err := t.Purchase.Validate(); err != nil {
		var v *app_errors.Validation
		if errors.As(err, &v) {
			errs = append(errs, v.Unwrap()...)
		}
	}

	if err := t.Store.Validate(); err != nil {
		var v *app_errors.Validation
		if errors.As(err, &v) {
			errs = append(errs, v.Unwrap()...)
		}
	}

	if err := t.Acquirer.Validate(); err != nil {
		var v *app_errors.Validation
		if errors.As(err, &v) {
			errs = append(errs, v.Unwrap()...)
		}
	}

	if len(errs) > 0 {
		return app_errors.NewValidation(errs...)
	}

	return nil
}
