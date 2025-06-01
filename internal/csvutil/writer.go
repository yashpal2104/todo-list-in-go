package csvutil

import (
	"encoding/csv"
	"fmt"

	
)

func writeCSVRecord(filepath string, record []string) error {
	f, err := loadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to load file: %w", err)
	}
	defer closeFile(f)

	writer := csv.NewWriter(f)
	defer writer.Flush()

	if err := writer.Write(record); err != nil {
		return fmt.Errorf("failed to write record: %w", err)
	}

	return nil
}
