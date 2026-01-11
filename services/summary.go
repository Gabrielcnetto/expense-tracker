package services

import "fmt"

func Summary() (float32, error) {
	expenses, err := GetAll()
	if err != nil {
		return 0, err
	}
	var total float32
	for _, item := range expenses {
		fmt.Println(item.Amount)
		total += item.Amount
	}

	return total, nil
}
