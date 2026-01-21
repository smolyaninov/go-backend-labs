package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-backend-labs/01-beginner/expense-tracker/internal/service"
	"log"
	"time"
)

var (
	budgetMonth  int
	budgetAmount float64
)

var setBudgetCmd = &cobra.Command{
	Use:   "set-budget",
	Short: "Set a budget for a month",
	Run: func(cmd *cobra.Command, args []string) {
		if budgetMonth < 1 || budgetMonth > 12 {
			fmt.Println("Month must be between 1 and 12")
		}
		if budgetAmount <= 0 {
			fmt.Println("Amount must be greater than zero")
		}

		budgetService := service.NewDefaultBudgetService()

		year := time.Now().Year()
		if err := budgetService.SetBudget(budgetMonth, year, budgetAmount); err != nil {
			log.Fatalf("Error setting budget: %v", err)
		}

		fmt.Printf("Budget for month %d set to $%.2f\n", budgetMonth, budgetAmount)
	},
}

func init() {
	rootCmd.AddCommand(setBudgetCmd)

	setBudgetCmd.Flags().IntVarP(&budgetMonth, "month", "m", 0, "Month (1-12)")
	setBudgetCmd.Flags().Float64VarP(&budgetAmount, "amount", "a", 0, "Budget amount")

	setBudgetCmd.MarkFlagRequired("month")
	setBudgetCmd.MarkFlagRequired("amount")
}
