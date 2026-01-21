package domain

import (
	"fmt"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
}

func NewExpense(id int, description string, amount float64, category string) (*Expense, error) {
	if description == "" {
		return nil, fmt.Errorf("description cannot be empty")
	}

	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Date:        time.Now().UTC(),
		Category:    category,
	}, nil
}

func (e *Expense) Update(description string, amount float64, category string) error {
	if description != "" {
		e.Description = description
	}

	if amount > 0 {
		e.Amount = amount
	} else if amount < 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	if category != "" {
		e.Category = category
	}

	return nil
}

func (e *Expense) BelongsToMonth(month int, year int) bool {
	return e.Date.Month() == time.Month(month) && e.Date.Year() == year
}
