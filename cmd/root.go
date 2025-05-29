/*
Copyright © 2025 Yash  <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "It helps you select fruits of different sizes and test out the Renovate tool",
	Long:  `This CLI is designed for you to select fruits based on different sizes blah blah blah you don't care about this even I don't care about it what we are trying to do here is to test out the Renovate tool to automate dependency updates and patch regular security threats through making PRs to fix it.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {

	// },
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all the fruits",
	Run: func(cmd *cobra.Command, args []string) {
		for i, description := range data.FileParsedData {
			fmt.Println("%d. %+v\n", description, i+1, description)
		}
	},
}

var FilterCmd = &cobra.Command{
	Use:   "filter",
	Short: "filters the fruits",
	Run: func(cmd *cobra.Command, args []string) {
		for i, description := range data.FileParsedData {
			// fmt.Println("%d. %+v\n", description, i+1, description)
			fmt.Println("Fruit: %s Color: %s", description.Fruit, description.Color, i+1)

		}
	},
}

var colorFlag string
var sizeFlag string

var Color = &cobra.Command{
	Use:   "filter",
	Short: "filters the fruits",
	Run: func(cmd *cobra.Command, args []string) {
		for i, description := range data.FileParsedData {
			// fmt.Println("%d. %+v\n", description, i+1, description)
			// Example: filter by color passed as the first argument after the command
			if len(args) > 0 && description.Color == args[0] {
				colorFlag = description.Color
				fmt.Println("Fruit: %s Color: %s\n", description.Fruit, description.Color, i+1)
			}
		}
	},
}

var Size = &cobra.Command{
	Use:   "size",
	Short: "filters the fruits based on size",
	Run: func(cmd *cobra.Command, args []string) {
		for i, description := range data.FileParsedData {
			// fmt.Println("%d. %+v\n", description, i+1, description)
			// Example: filter by color passed as the first argument after the command
			if len(args) > 0 && description.Size == args[0] {
				sizeFlag = description.Size
				fmt.Println("Fruit: %s Size: %s\n", description.Fruit, description.Size, i+1)
			}
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.json-viewer-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(FilterCmd)
	FilterCmd.Flags().StringVarP(&colorFlag, "--color", "c", "", "Filter fruits by color")
	FilterCmd.Flags().StringVarP(&sizeFlag, "--size", "s", "", "Filter fruits by size")
}
