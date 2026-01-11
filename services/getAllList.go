package services

import "github.com/Gabrielcnetto/expense-tracker/models"

func GetAll() ([]models.Expense, error) {
	allTasks, err := GetExpenses()
	return allTasks, err
}
