/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Gabrielcnetto/expense-tracker/services"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove expense of list",
	Long:  `Update list removing one expense item`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			return err
		}
		err = services.Delete(id)
		if err != nil {
			return err
		}
		fmt.Printf("\nExpense (ID %v) deleted successfully\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int("id", 0, "Expense id for delete from list")
}
