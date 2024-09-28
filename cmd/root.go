/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gemix",
	Short: "API endpoint fuzz testing with Gemini",
	Long: `Gemix is a powerful tool that integrates with Gemini to perform fuzz testing on your API endpoints.

It generates a wide range of test cases to thoroughly examine the robustness and security of your API.
By leveraging Gemini's capabilities, Gemix helps identify potential vulnerabilities, edge cases,
and unexpected behaviors in your API endpoints.

Use Gemix to enhance the quality and reliability of your API by subjecting it to comprehensive fuzz testing.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Main logic for running the fuzz tests
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gemix.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
