/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thenakulchawla/refactored-waddle/pkg/arrays"
)

// arraysCmd represents the arrays command
var arraysCmd = &cobra.Command{
	Use:   "arrays",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// change this to test the function you want to run
		arrays.LongCons()
	},
}

func init() {
	rootCmd.AddCommand(arraysCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// arraysCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// arraysCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
