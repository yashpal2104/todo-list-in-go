package cmd

import (
	"encoding/csv"
	"os"
	"strconv"

)

// func WriteCSVRecord(filepath string, records [][]string) error {
// 	f, err := loadFile(filepath)
// 	if err != nil {
// 		return fmt.Errorf("failed to load file: %w", err)
// 	}
// 	defer closeFile(f)

// 	writer := csv.NewWriter(f)
// 	defer writer.Flush()

// 	return writer.WriteAll(records)
// }

func AppendCSVRecord(filepath string, item Item) error{
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    record := []string{
        strconv.Itoa(item.ID),
        item.Description,
        HumanizeTimeSince(item.CreatedAt),
    }
    return writer.Write(record)
}
	