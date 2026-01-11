package services

import "github.com/Gabrielcnetto/expense-tracker/models"

func AddNewExpense(description string, amount float32) (int, error) {
	newExpense := models.Expense{
		Description: description,
		Amount:      amount,
	}
	newExpense.ConvertAmount()
	newExpense.GenDate()
	id, response := Manager(newExpense)
	return id, response
}
