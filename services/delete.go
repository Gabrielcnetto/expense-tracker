package services

import "github.com/Gabrielcnetto/expense-tracker/models"

func Delete(id int) error {
	allTasks, err := GetAll()
	if err != nil {
		return err
	}
	updatedList := []models.Expense{}
	for _, item := range allTasks {
		if item.Id != int8(id) {
			updatedList = append(updatedList, item)
		}
		continue
	}
	err = Save(updatedList)
	return err
}
