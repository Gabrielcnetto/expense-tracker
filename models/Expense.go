package models

import (
	"fmt"
	"time"
)

type Expense struct {
	Id          int8      `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	FmtAmount   string    `json:"fmt_amount"`
}

func (e *Expense) GenDate() {
	e.Date = time.Now()
}

func (e *Expense) ConvertAmount() {
	amount := fmt.Sprintf("$%.2f", e.Amount)
	e.FmtAmount = amount
}
