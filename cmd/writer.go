package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func WriteAllCSVRecord(filepath string, records [][]string) error {
	f, err := loadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to load file: %w", err)
	}
	defer closeFile(f)

	writer := csv.NewWriter(f)
	defer writer.Flush()

	return writer.WriteAll(records)
}

func AppendCSVRecord(filepath string, item Item) error {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header if file is new/empty
	info, err := file.Stat()
	if err == nil && info.Size() == 0 {
		if err := writer.Write([]string{"ID", "Description", "CreatedAt"}); err != nil {
			return err
		}
	}

	record := []string{
		strconv.Itoa(item.ID),
		item.Description,
		item.CreatedAt.Format(time.RFC3339), // Store the real timestamp!
	}
	return writer.Write(record)
}

// CheckFileIsExist checks if a file exists at the given path.
func CheckFileIsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}