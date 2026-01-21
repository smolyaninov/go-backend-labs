package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-backend-labs/01-beginner/expense-tracker/internal/service"
	"log"
	"time"
)

var (
	description string
	amount      float64
	category    string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		expenseService := service.NewDefaultExpenseService()

		expense, err := expenseService.AddExpense(description, amount, category)
		if err != nil {
			log.Fatalf("Error adding expense: %v", err)
		}

		fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)

		budgetService := service.NewDefaultBudgetService()

		year := time.Now().Year()
		month := time.Now().Month()

		budget, err := budgetService.GetBudget(int(month), year)
		if err != nil {
			log.Printf("Warning: could not load budget: %v", err)
		}

		if budget > 0 {
			fmt.Printf("\nBudget for %s: $%.2f\n", month, budget)

			total, err := expenseService.GetTotalAmountFiltered(int(month), "")
			if err == nil {
				fmt.Printf("Current total for %s: $%.2f\n", month, total)

				if total > budget {
					fmt.Printf("Warning: you have exceeded your budget for %s!\n", month)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0, "Amount")
	addCmd.Flags().StringVarP(&category, "category", "c", "", "Category")

	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
