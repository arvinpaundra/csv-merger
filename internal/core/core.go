package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/arvinpaundra/csv-merger/internal/csv"
)

func validateFile(filename string) error {
	ext := filepath.Ext(filename)
	if ext != ".csv" {
		return fmt.Errorf("invalid file type")
	}
	return nil
}

func ReadCsv(path string) (csv.Records, error) {
	LogInfo("start reading ", path)

	defer LogSuccess("done")

	err := validateFile(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var csv csv.Records

	err = csv.Read(f)
	if err != nil {
		return nil, err
	}

	return csv, nil
}

func FindKeyIndex(key string, csv csv.Records) (n int, err error) {
	LogInfo("finding key ", key)

	if len(csv) > 0 {
		for i := range csv[0] {
			if csv[0][i] == key {
				LogSuccess("done")

				return i, nil
			}
		}
	}

	return -1, fmt.Errorf("missing key %s", key)
}

func MergeCsv(keyIndex1, keyIndex2 int, records1, records2 csv.Records) (csv.Records, error) {
	LogInfo("merging file")

	if len(records1) == 0 || len(records2) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	defer LogSuccess("done")

	if len(records1) > len(records2) {
		return merge(keyIndex2, keyIndex1, records2, records1), nil
	}

	return merge(keyIndex1, keyIndex2, records1, records2), nil
}

func merge(keyIndex1, keyIndex2 int, records1, records2 csv.Records) csv.Records {
	recordsMap := make(map[string][]string)

	for _, record := range records1[1:] {
		key := record[keyIndex1]
		recordsMap[key] = record
	}

	header1 := records1[0]
	header2 := records2[0]

	mergeHeaders := append(header1, header2[:keyIndex2]...)
	mergeHeaders = append(mergeHeaders, header2[keyIndex2+1:]...)

	results := csv.Records{mergeHeaders}

	for _, record := range records2[1:] {
		key := record[keyIndex2]
		if record2, exist := recordsMap[key]; exist {
			merged := append(record2, record[:keyIndex2]...)
			merged = append(merged, record[keyIndex2+1:]...)

			results = append(results, merged)
		}
	}

	return results
}

func WriteCsv(filename string, records csv.Records) error {
	LogInfo("write csv into disk")

	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer out.Close()

	err = records.Write(out)
	if err != nil {
		return err
	}

	LogSuccess("done")

	return nil
}
