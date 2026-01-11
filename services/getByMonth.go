package services

import (
	"time"
)

func GetByMonth(month int) (float32, error) {
	expenses, err := GetAll()
	if err != nil {
		return 0, err
	}

	var total float32
	for _, item := range expenses {
		if item.Date.Month() == time.Month(month) {
			total += item.Amount
		}
		continue
	}
	return total, nil
}
