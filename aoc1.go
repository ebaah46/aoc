package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var totalCalibrationScore = 0

func GenCalibrationValue(line string) (int, error) {
	regex := regexp.MustCompile(`\d`)
	firstDigit := regex.FindString(line)
	regex = regexp.MustCompile("(\\d)\\D*$")
	rawEnd := regex.FindString(line)
	lastDigit := string(rawEnd[0]) // workaround for go not having raw +/- lookahead in regex
	combinedDigits := firstDigit + lastDigit
	calibrationScore, err := strconv.Atoi(combinedDigits)
	if err != nil {
		return 0, err
	}
	return calibrationScore, nil
}

func main() {
	pathToFile := "data/d1"
	data, err := ioutil.ReadFile(pathToFile)

	if err != nil {
		fmt.Printf("Read failed with: %s", err.Error())
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		calibrationValue, err := GenCalibrationValue(line)
		if err == nil {
			totalCalibrationScore += calibrationValue
		}
	}
	fmt.Printf("Total calibration value is: %d", totalCalibrationScore)
}

// Algo for part 2
/*
	1. Check the length of each string.
	2. Check for digit positions and store indexes of digits.
	3. Check for lettered digits.
*/
