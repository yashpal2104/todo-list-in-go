package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"

)

func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func appendTaskToCSV(filename string, item Item) error {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
