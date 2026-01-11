/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Gabrielcnetto/expense-tracker/services"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Get a sum of all expenses that you added",
	Long:  `Your Total expenses`,
	RunE: func(cmd *cobra.Command, args []string) error {
		month, _ := cmd.Flags().GetInt("month")
		if month == 0 {
			total, err := services.Summary()
			if err != nil {
				return err
			}
			fmt.Printf("\nTotal expenses: $%.2f \n", total)
			return nil
		}
		total, err := services.GetByMonth(month)
		if err != nil {
			return err
		}
		fmt.Printf("\nTotal expenses: $%.2f \n", total)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
	summaryCmd.Flags().Int("month", 0, "Month that you wanna get")
}
