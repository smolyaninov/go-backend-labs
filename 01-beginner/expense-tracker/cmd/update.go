package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-backend-labs/01-beginner/expense-tracker/internal/service"
	"log"
)

var (
	updateID          int
	updateDescription string
	updateAmount      float64
	updateCategory    string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an expense by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if updateDescription == "" && updateAmount == 0 && updateCategory == "" {
			log.Fatalf("You must provide at least one field to update: --description, --amount, --category\n")
		}

		expenseService := service.NewDefaultExpenseService()

		err := expenseService.UpdateExpense(updateID, updateDescription, updateAmount, updateCategory)
		if err != nil {
			log.Fatalf("Error updating expense: %v", err)
		}

		fmt.Println("Expense updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateID, "id", "i", 0, "ID of the expense to update")
	updateCmd.Flags().StringVarP(&updateDescription, "description", "d", "", "New description")
	updateCmd.Flags().Float64VarP(&updateAmount, "amount", "a", 0, "New amount")
	updateCmd.Flags().StringVarP(&updateCategory, "category", "c", "", "New category")

	updateCmd.MarkFlagRequired("id")
}
