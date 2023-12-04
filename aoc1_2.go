package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func GenCalibrationValue2(line string) (int, error) {

	return 0, nil
}

func main() {
	pathToFile := "data/d1"
	data, err := ioutil.ReadFile(pathToFile)

	if err != nil {
		fmt.Printf("Read failed with: %s", err.Error())
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		calibrationValue, err := GenCalibrationValue2(line)
		if err == nil {
			totalCalibrationScore += calibrationValue
		}
	}
	fmt.Printf("Total calibration value is: %d", totalCalibrationScore)

}
