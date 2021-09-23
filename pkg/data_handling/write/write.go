package data_handling

import (
	"os"
	"strings"

	"github.com/pkg/errors"
)

// Writes ciks to a given file
func WriteFile(cik_lines []string, path string) (err error) {
	var (
		sec_file *os.File
	)

	if sec_file, err = os.Create(path); err != nil {
		errors.Errorf("Error truncating file %s", sec_file)
	}
	defer sec_file.Close()

	for _, file_item := range cik_lines {
		_, err := sec_file.WriteString(strings.TrimSpace(file_item) + "\n")
		sec_file.Write([]byte(file_item))
		if err != nil {
			errors.Errorf("Error writing file item %s as string", file_item)
			break
		}
	}
	return
}
