package data_handling

import (
	"bufio"
	"bytes"
	"io"
	"os"

	"github.com/pkg/errors"
)

// Reads file of CIKs from SEC text file and stores as an array
func ReadFile(path string) (cik_lines []string, err error) {
	var (
		sec_file     *os.File
		part_of_file []byte
		prefix_check bool
	)

	if sec_file, err = os.Open(path); err != nil {
		errors.Errorf("Error reading %s on path %s", sec_file, path)
	}
	defer sec_file.Close()

	reader := bufio.NewReader(sec_file)
	buffer := bytes.NewBuffer(make([]byte, 0))

	for {
		if part_of_file, prefix_check, err = reader.ReadLine(); err != nil {
			errors.Errorf("Error reading %s", part_of_file)
			break
		}
		buffer.Write(part_of_file)
		if !prefix_check {
			cik_lines = append(cik_lines, buffer.String())
			buffer.Reset()
		}
	}

	if err == io.EOF {
		err = nil
	}
	return
}
