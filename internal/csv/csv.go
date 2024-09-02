package csv

import (
	gocsv "encoding/csv"
	"fmt"
	"io"
)

type Records [][]string

func (records *Records) Read(r io.Reader) error {
	reader := gocsv.NewReader(r)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		*records = append(*records, record)
	}
}

func (records *Records) Write(w io.Writer) error {
	if records == nil {
		return fmt.Errorf("records empty")
	}

	writer := gocsv.NewWriter(w)

	for _, record := range *records {
		err := writer.Write(record)

		if err != nil {
			return err
		}
	}

	writer.Flush()

	return writer.Error()
}
