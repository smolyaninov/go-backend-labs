package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-backend-labs/01-beginner/expense-tracker/internal/domain"
	"go-backend-labs/01-beginner/expense-tracker/internal/service"
	"log"
	"os"
	"text/tabwriter"
)

var listCategory string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		expenseService := service.NewDefaultExpenseService()

		expenses, err := expenseService.GetAllExpenses()
		if err != nil {
			log.Fatalf("Error getting expenses: %v", err)
		}

		if len(expenses) == 0 {
			fmt.Println("No expenses found")
			return
		}

		if listCategory != "" {
			newList := make([]domain.Expense, 0)
			for _, e := range expenses {
				if e.Category == listCategory {
					newList = append(newList, e)
				}
			}
			expenses = newList

			if len(expenses) == 0 {
				fmt.Printf("No expenses found for category %s\n", listCategory)
				return
			}
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(w, "ID\tDate\tDescription\tAmount\tCategory")
		for _, e := range expenses {
			fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\t%s\n", e.ID, e.Date.Format("2006-01-02"), e.Description, e.Amount, e.Category)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&listCategory, "category", "c", "", "Filter expenses by category")
}
