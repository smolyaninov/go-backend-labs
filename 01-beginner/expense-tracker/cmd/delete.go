package cmd

import (
	"fmt"
	"go-backend-labs/01-beginner/expense-tracker/internal/service"

	"github.com/spf13/cobra"
)

var deleteID int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense by ID",
	Run: func(cmd *cobra.Command, args []string) {
		expenseService := service.NewDefaultExpenseService()

		err := expenseService.DeleteExpense(deleteID)
		if err != nil {
			fmt.Printf("Error deleting expense: %v", err)
			return
		}

		fmt.Println("Expense deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&deleteID, "id", "i", 0, "ID")
	deleteCmd.MarkFlagRequired("id")
}
