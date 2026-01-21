package service

import "go-backend-labs/01-beginner/expense-tracker/internal/repo"

func NewDefaultExpenseService() *ExpenseService {
	return NewExpenseService(repo.NewJSONExpenseRepository("data/expense.json"))
}

func NewDefaultBudgetService() *BudgetService {
	return NewBudgetService(repo.NewJSONBudgetRepository("data/budget.json"))
}
