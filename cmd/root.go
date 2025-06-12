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
			fmt.Fprintf(w, "%d\t%s\t%s\n", item.ID, item.Description, timediff.TimeDiff(item.CreatedAt))
		}
		w.Flush()
		// fmt.Fprintln(w, "ID\tDescription\tCreatedAt")

	},
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add all the tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// if len(Args) < 2 {
		// 	fmt.Println("Expected 'add' command")
		// 	return
		// }
		// switch Args[1] {
		// case "add":
		// 	if len(Args) < 3 {
		// 		fmt.Println("example usage: tasks add <description>")
		// 		return
		// 	}
		// }
		// description := Args[2]
		// fmt.Println("Adding task:", description)

		if len(args) == 0 {
            fmt.Println("example usage: tasks add <description1> <description2> ...")
            return
        }

		
		// for i, description := range os.Args[2:] {
		// data = append(data, Item{ID: i + 1, Description: description, CreatedAt: time.Since(start)})

		// fmt.Println("%d. %+v\n", description, i+1, description)
		// }
		for _, description := range args{
			newItem := Item{
			ID: len(data) + 1,
			Description: description,
			CreatedAt: time.Now(),
		}
		data = append(data, newItem)
		fmt.Println("Adding task:", description)
		}
		
		// data = append(data, newItem)

		
		records := [][]string{{"ID", "Description", "CreatedAt"}}
		for _, item := range data {
			fmt.Println(strconv.Itoa(item.ID) + "." + item.Description + " " + item.CreatedAt.String())
			records = append(records, []string{
				strconv.Itoa(item.ID), item.Description, HumanizeTimeSince(item.CreatedAt),
			})
		}
		// for _, record := range records{
			err := AppendCSVRecord("output.csv", records)
			if err != nil {
			log.Fatalf("error writing CSV: %v", err)
		}
		// }
		
		// fmt.Println(records[1][0])
	},
}

// var FilterCmd = &cobra.Command{
// 	Use:   "filter",
// 	Short: "filters the fruits",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		for i, description := range data.FileParsedData {
// 			// fmt.Println("%d. %+v\n", description, i+1, description)
// 			fmt.Println("Fruit: %s Color: %s", description.Fruit, description.Color, i+1)

// 		}
// 	},
// }

// var colorFlag string
// var sizeFlag string

// var Color = &cobra.Command{
// 	Use:   "filter",
// 	Short: "filters the fruits",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		for i, description := range data.FileParsedData {
// 			// fmt.Println("%d. %+v\n", description, i+1, description)
// 			// Example: filter by color passed as the first argument after the command
// 			if len(args) > 0 && description.Color == args[0] {
// 				colorFlag = description.Color
// 				fmt.Println("Fruit: %s Color: %s\n", description.Fruit, description.Color, i+1)
// 			}
// 		}
// 	},
// }

// var Size = &cobra.Command{
// 	Use:   "size",
// 	Short: "filters the fruits based on size",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		for i, description := range data.FileParsedData {
// 			// fmt.Println("%d. %+v\n", description, i+1, description)
// 			// Example: filter by color passed as the first argument after the command
// 			if len(args) > 0 && description.Size == args[0] {
// 				sizeFlag = description.Size
// 				fmt.Println("Fruit: %s Size: %s\n", description.Fruit, description.Size, i+1)
// 			}
// 		}
// 	},
// }

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	w = tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(AddCmd)
	// abs, _ := filepath.Abs(csvFilePath)
	// fmt.Println("Checking file at:", abs)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.json-viewer-cli.yaml)")
	rootCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.AddCommand(ListCmd)
	// rootCmd.AddCommand(FilterCmd)
	// FilterCmd.Flags().StringVarP(&colorFlag, "--color", "c", "", "Filter fruits by color")
	// FilterCmd.Flags().StringVarP(&sizeFlag, "--size", "s", "", "Filter fruits by size")
}
