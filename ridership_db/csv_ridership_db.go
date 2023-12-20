package ridershipDB

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CsvRidershipDB struct {
	idIdxMap      map[string]int
	csvFile       *os.File
	csvReader     *csv.Reader
	num_intervals int
}

func (c *CsvRidershipDB) Open(filePath string) error {
	c.num_intervals = 9

	// Create a map that maps MBTA's time period ids to indexes in the slice
	c.idIdxMap = make(map[string]int)
	for i := 1; i <= c.num_intervals; i++ {
		timePeriodID := fmt.Sprintf("time_period_%02d", i)
		c.idIdxMap[timePeriodID] = i - 1
	}

	// create csv reader
	csvFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	c.csvFile = csvFile
	c.csvReader = csv.NewReader(c.csvFile)

	return nil
}

// TODO: some code goes here
// Implement the remaining RidershipDB methods
func (c *CsvRidershipDB) GetRidership(lineId string) ([]int64, error) {
	// return res
	var res = make([]int64, 9)
	for {
		record, err := c.csvReader.Read()
		// if record == nil, there is no more record, break
		if record == nil {
			break
		}
		if err != nil {
			panic(err.Error())
		}
		// sum the ons
		if record[0] == lineId {
			// time_period_0i -> idx = i-1
			idx := c.idIdxMap[record[2]]
			ons, _ := strconv.Atoi(record[4])
			res[idx] += int64(ons)
		}

	}
	return res, nil
}

func (c *CsvRidershipDB) Close() error {
	return nil
}
