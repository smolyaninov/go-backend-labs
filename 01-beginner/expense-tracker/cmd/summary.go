package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-backend-labs/01-beginner/expense-tracker/internal/service"
	"log"
)

var (
	summaryMonth    int
	summaryCategory string
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show total expenses (optional by month)",
	Run: func(cmd *cobra.Command, args []string) {
		expenseService := service.NewDefaultExpenseService()

		if summaryMonth != 0 && (summaryMonth < 1 || summaryMonth > 12) {
			log.Fatalf("Month must be between 1 and 12")
		}

		total, err := expenseService.GetTotalAmountFiltered(summaryMonth, summaryCategory)
		if err != nil {
			log.Fatalf("Error getting total expenses: %v", err)
		}

		switch {
		case summaryMonth > 0 && summaryCategory != "":
			fmt.Printf("Total expenses for month %d and category %s: $%.2f\n", summaryMonth, summaryCategory, total)
		case summaryMonth > 0:
			fmt.Printf("Total expenses for month %d: $%.2f\n", summaryMonth, total)
		case summaryCategory != "":
			fmt.Printf("Total expenses for category %s: $%.2f\n", summaryCategory, total)
		default:
			fmt.Printf("Total expenses: $%.2f\n", total)
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntVarP(&summaryMonth, "month", "m", 0, "Filter summary by month (1-12)")
	summaryCmd.Flags().StringVarP(&summaryCategory, "category", "c", "", "Filter summary by category")
}
