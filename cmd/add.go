/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Gabrielcnetto/expense-tracker/services"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new Expense to list",
	Long:  `Add using this command new expense for general list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		description, err := cmd.Flags().GetString("description")
		if err != nil {
			return err
		}
		amount, err := cmd.Flags().GetFloat32("amount")
		if err != nil {
			return err
		}
		id, err := services.AddNewExpense(description, amount)
		if err != nil {
			return err
		}
		fmt.Printf("\nExpense added successfully (ID %v)\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("description", "", "Description for your expense")
	addCmd.Flags().Float32("amount", 0, "Amount for your expense")
}
