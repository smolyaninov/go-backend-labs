package service

import (
	"fmt"
	"go-backend-labs/01-beginner/expense-tracker/internal/repo"
)

type BudgetService struct {
	repo repo.BudgetRepository
}

func NewBudgetService(repo repo.BudgetRepository) *BudgetService {
	return &BudgetService{repo}
}

func (s *BudgetService) SetBudget(month int, year int, amount float64) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("month must be between 1 and 12")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}

	key := fmt.Sprintf("%d-%d", month, year)
	budgets, err := s.repo.Load()
	if err != nil {
		return err
	}

	if budgets == nil {
		budgets = make(map[string]float64)
	}

	budgets[key] = amount
	return s.repo.Save(budgets)
}

func (s *BudgetService) GetBudget(month int, year int) (float64, error) {
	key := fmt.Sprintf("%d-%d", month, year)
	budgets, err := s.repo.Load()
	if err != nil {
		return 0, err
	}

	return budgets[key], nil
}
