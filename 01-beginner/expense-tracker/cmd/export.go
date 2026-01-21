package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"go-backend-labs/01-beginner/expense-tracker/internal/service"
	"log"
	"os"
	"time"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export expenses to CSV",
	Run: func(cmd *cobra.Command, args []string) {
		expenseService := service.NewDefaultExpenseService()

		expenses, err := expenseService.GetAllExpenses()
		if err != nil {
			log.Fatalf("Error getting expenses: %v", err)
		}

		if len(expenses) == 0 {
			fmt.Println("No expenses to export")
			return
		}

		exportDir := "data/exports"
		if err := os.MkdirAll(exportDir, 0755); err != nil {
			log.Fatalf("Error creating directories: %v", err)
		}

		timestamp := time.Now().Format("2006-01-02_15-04-05")
		fileName := fmt.Sprintf("export_%s.csv", timestamp)
		fullPath := fmt.Sprintf("%s/%s", exportDir, fileName)

		file, err := os.Create(fullPath)
		if err != nil {
			log.Fatalf("Error creating file: %v", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		err = writer.Write([]string{"ID", "Date", "Description", "Amount", "Category"})
		if err != nil {
			log.Fatalf("Error writing header: %v", err)
		}

		for _, e := range expenses {
			row := []string{
				fmt.Sprintf("%d", e.ID),
				e.Date.Format("2006-01-02"),
				e.Description,
				fmt.Sprintf("%.2f", e.Amount),
				e.Category,
			}
			if err := writer.Write(row); err != nil {
				log.Fatalf("Error writing row: %v", err)
			}
		}

		fmt.Printf("Exported %d expenses to %s\n", len(expenses), fullPath)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
