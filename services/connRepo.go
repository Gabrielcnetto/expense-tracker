package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Gabrielcnetto/expense-tracker/models"
)

const (
	REPO = "./repositories/data.json"
)

func verifyRepo() int {
	dir := filepath.Dir(REPO)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		//create repo
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return 0
		}
	}
	if _, err := os.Stat(REPO); os.IsNotExist(err) {
		//create file
		file, err := os.Create(REPO)
		if err != nil {
			return 0
		}
		defer file.Close()
		return 1
	}
	return 1
}

func Manager(item models.Expense) (int, error) {

	verify := verifyRepo()
	switch verify {
	case 1:
		allTasks, _ := GetExpenses()
		id := len(allTasks) + 1
		item.Id = int8(id)
		allTasks = append(allTasks, item)
		responseSave := Save(allTasks)
		return id, responseSave
	case 2:
		return 0, fmt.Errorf("Error to save on repo")
	}
	return 0, nil

}
func Save(items []models.Expense) error {

	data, err := json.MarshalIndent(items, "", "")
	if err != nil {
		return err
	}
	err = os.WriteFile(REPO, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
func GetExpenses() ([]models.Expense, error) {
	data, err := os.ReadFile(REPO)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("Actual data is empty")
	}
	var tasks []models.Expense
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
