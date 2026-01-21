package service

import (
	"fmt"
	"go-backend-labs/01-beginner/expense-tracker/internal/domain"
	"go-backend-labs/01-beginner/expense-tracker/internal/repo"
	"time"
)

type ExpenseService struct {
	repo repo.ExpenseRepository
}

func NewExpenseService(repo repo.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo}
}

func (s *ExpenseService) AddExpense(description string, amount float64, category string) (domain.Expense, error) {
	expenses, err := s.repo.Load()
	if err != nil {
		return domain.Expense{}, err
	}

	newID := 1
	if len(expenses) > 0 {
		newID = expenses[len(expenses)-1].ID + 1
	}

	expense, err := domain.NewExpense(newID, description, amount, category)
	if err != nil {
		return domain.Expense{}, err
	}

	expenses = append(expenses, *expense)

	if err := s.repo.Save(expenses); err != nil {
		return domain.Expense{}, err
	}

	return *expense, nil
}

func (s *ExpenseService) GetAllExpenses() ([]domain.Expense, error) {
	return s.repo.Load()
}

func (s *ExpenseService) DeleteExpense(id int) error {
	expenses, err := s.repo.Load()
	if err != nil {
		return err
	}

	newExpenses := make([]domain.Expense, 0)
	found := false

	for _, e := range expenses {
		if e.ID == id {
			found = true
			continue
		}
		newExpenses = append(newExpenses, e)
	}

	if !found {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	return s.repo.Save(newExpenses)
}

func (s *ExpenseService) UpdateExpense(id int, newDescription string, newAmount float64, newCategory string) error {
	expenses, err := s.repo.Load()
	if err != nil {
		return err
	}

	updated := false

	for i := range expenses {
		if expenses[i].ID == id {
			if err := expenses[i].Update(newDescription, newAmount, newCategory); err != nil {
				return err
			}
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	return s.repo.Save(expenses)
}

func (s *ExpenseService) GetTotalAmountFiltered(month int, category string) (float64, error) {
	if month != 0 && (month < 1 || month > 12) {
		return 0, fmt.Errorf("month must be between 1 and 12")
	}

	expenses, err := s.repo.Load()
	if err != nil {
		return 0, err
	}

	currentYear := time.Now().Year()
	var total float64

	for _, e := range expenses {
		if month != 0 && !e.BelongsToMonth(month, currentYear) {
			continue
		}

		if category != "" && e.Category != category {
			continue
		}

		total += e.Amount
	}

	return total, nil
}
