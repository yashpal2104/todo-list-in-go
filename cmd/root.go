/*
Copyright © 2025 Yash  <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	// "path/filepath"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
	// "github.com/mergestat/timediff"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

type Item struct {
	ID          int
	Description string
	CreatedAt   time.Time
}

var data []Item
var csvFilePath = "./output.csv"
var deleteAll bool
var w *tabwriter.Writer

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "TODO CLI App",
	Long:  `This CLI helps manage tasks`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all the tasks",
	Run: func(cmd *cobra.Command, args []string) {

		if !CheckFileIsExist(csvFilePath) {
			fmt.Println("There are currently no tasks. Please use the 'tasks add <description>' to add your tasks")

			return
		}

		taskList, err := ReadAndWriteCSVTasks(csvFilePath)
		if err != nil {
			fmt.Errorf("Trouble loading the tasks list: %v\n", err)
			return
		}
		if len(taskList) == 0 {
			fmt.Println("There are currently no tasks. Please use the 'tasks add <description>' to add your tasks")
			return
		}
		// for _, item := range taskList {
		// 	fmt.Printf("%d. %s (%s)\n", item.ID, item.Description, HumanizeTimeSince(item.CreatedAt))
		// }
		fmt.Fprintf(w, "ID\tDescription\tCreatedAt\n")
		for _, item := range taskList {
			fmt.Fprintf(w, "%d\t%s\t%s\t\n", item.ID, item.Description, timediff.TimeDiff(item.CreatedAt))
		}
		w.Flush()
		// fmt.Fprintln(w, "ID\tDescription\tCreatedAt")

	},
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add all the tasks",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("example usage: tasks add <description1> <description2> ...")
			return
		}
		existingTasks, _ := ReadAndWriteCSVTasks(csvFilePath)
		nextID := getLastID(existingTasks) + 1
		for _, description := range args {
			newItem := Item{
				ID:          nextID,
				Description: description,
				CreatedAt:   time.Now(),
			}
			nextID++
			data = append(data, newItem)
			err := AppendCSVRecord(csvFilePath, newItem)
			if err != nil {
				log.Fatalf("error writing CSV: %v", err)
			}
			fmt.Println("Added task: ", newItem.Description)
		}
		w.Flush()

	},
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete the tasks specified in the args",
	Run: func(cmd *cobra.Command, args []string) {
		if !deleteAll && len(args) == 0 {
			fmt.Println("example usage: tasks delete <description1> <description2> ...")
			return
		}
		if deleteAll {
            // Delete all tasks
            data = []Item{}
            err := BuildRecordsForCSV()
            if err != nil {
                log.Fatal("error deleting all tasks from CSV: ", err)
            }
            fmt.Println("All tasks deleted.\n")
            ListCmd.Run(cmd, []string{})
            w.Flush()
            return
        }
		_, err := DeleteTasksFromCSV(csvFilePath, args)
		if err != nil {
			log.Fatal("error deleting the tasks from CSV: ", err)
		}

		for _, desc := range args {
			// desc is each description passed as an argument
			fmt.Println("Deleted task: ", desc)
		}
		ListCmd.Run(cmd, []string{})
		w.Flush()
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(DeleteCmd)
	// abs, _ := filepath.Abs(csvFilePath)
	// fmt.Println("Checking file at:", abs)

	rootCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")
	DeleteCmd.Flags().BoolVarP(&deleteAll, "all", "a", false, "Delete all tasks")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
