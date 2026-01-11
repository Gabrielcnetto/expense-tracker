/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/Gabrielcnetto/expense-tracker/services"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all expenses added by you",
	Long:  `A Complete expense list added by you via CLI commands  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		response, err := services.GetAll()
		if err != nil {
			return err
		}
		fmt.Printf("%-4s %-12s %-15s %10s\n", "ID", "DATE", "DESCRIPTION", "AMOUNT")
		fmt.Println(strings.Repeat("-", 45))
		for _, item := range response {
			fmt.Printf("%-4v %-12v %-15v %10v\n", item.Id, item.Date.Format("2006-01-02"), item.Description, item.FmtAmount)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
