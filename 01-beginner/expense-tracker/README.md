# Expense Tracker CLI

This is my third project in Go, completed as part of
the [Expense Tracker project](https://roadmap.sh/projects/expense-tracker) on roadmap.sh.

The goal was to build a more advanced CLI application with persistent storage, filtering, structured domain logic, and
extensibility — while continuing to practice Go idioms, interfaces, and generics.

The application is a command-line tool for tracking personal expenses, setting monthly budgets, and exporting reports.

## Features

- Add new expenses with amount, category, and description
- Update and delete existing expenses
- List all expenses, optionally filtered by category
- Show summary totals by month and/or category
- Set and get monthly budget limits
- Warn when budget is exceeded
- Export all expenses to CSV
- Persistent storage in local JSON files

## Usage

```bash
expense-cli add --description="Coffee" --amount=3.50 --category=food
expense-cli list
expense-cli list --category=food
expense-cli update --id=1 --amount=4.00
expense-cli delete --id=1
expense-cli summary --month=8 --category=food
expense-cli set-budget --month=8 --amount=300
expense-cli export
```

## Build

```bash
go build -o expense-cli ./main.go
```
