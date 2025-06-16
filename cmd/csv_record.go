package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "path/filepath"
	"strconv"
	"time"
)

// BuildRecordsForCSV writes all tasks in 'data' to the CSV file, overwriting it.
func BuildRecordsForCSV() error {
	records := [][]string{{"ID", "Description", "CreatedAt"}}
	for _, item := range data {
		records = append(records, []string{
			strconv.Itoa(item.ID),
			item.Description,
			item.CreatedAt.Format(time.RFC3339),
		})
	}
	err := WriteAllCSVRecords(csvFilePath, records)
	if err != nil {
		log.Fatalf("error writing CSV: %v", err)
		return err
	}
	return nil
}

// WriteAllCSVRecords writes all records to the given CSV file path.
func WriteAllCSVRecords(filePath string, records [][]string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

// ReadAndWriteCSVTasks reads all tasks from the CSV file and returns them as []Item.
func ReadAndWriteCSVTasks(filepath string) ([]Item, error) {

	file, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Errorf("Not able to open the file for reading")
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	// fmt.Println(records)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var tasks []Item

	for i, record := range records {
		// Skip header
		if i == 0 && record[0] == "ID" {
			continue
		}
		if len(record) < 3 {
			continue
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue // skip rows that can't be parsed
		}
		createdAt, err := time.Parse(time.RFC3339, record[2])
		if err != nil {
			createdAt = time.Now()
		}
		tasks = append(tasks, Item{
			ID:          id,
			Description: record[1],
			CreatedAt:   createdAt,
		})
	}
	return tasks, nil
}

// Delete the tasks from the csv file
func DeleteTasksFromCSV(filepath string, args []string) ([]Item, error){
	tasks, err := ReadAndWriteCSVTasks(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var updatedTasks []Item
	for _, task := range tasks{
		toDelete := false
		for _, desc := range args{
			if task.Description == desc {
			// Skip (delete) this task
			toDelete = true
			break
			}
		}
		if toDelete {
			continue //skip (delete) this task
		}
		// Appends only the tasks that are not given in the args
		updatedTasks = append(updatedTasks, task)	
	}
	// Temporarily replace your global data slice if needed
	data = updatedTasks
	err = BuildRecordsForCSV()
	if err != nil {
		log.Fatal(err)
}
	return updatedTasks, nil
}

// getLastID returns the highest ID in the given tasks slice.
func getLastID(tasks []Item) int {
	lastID := 0
	for _, item := range tasks {
		if item.ID > lastID {
			lastID = item.ID
		}
	}
	return lastID
}

