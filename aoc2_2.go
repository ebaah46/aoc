package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var coloredCubeCount = map[string]int{"red": 0, "green": 0, "blue": 0}

func GetGameSetPower(line string) (int, bool) {
	setPower := 1
	game := strings.Split(line, ":")
	_, found := strings.CutPrefix(game[0], "Game ")
	if !found {
		return -1, false
	}
	//gameNumber, _ = strconv.Atoi(number)
	sets := strings.Split(game[1], ";")
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			cubeData := strings.Split(cube, " ")
			drawCount, _ := strconv.Atoi(cubeData[0])
			if drawCount > coloredCubeCount[cubeData[1]] {
				coloredCubeCount[cubeData[1]] = drawCount
			}
		}
	}
	for key, val := range coloredCubeCount {
		setPower *= val
		// reset map values
		coloredCubeCount[key] = 0
	}

	return setPower, true
}

func main() {
	pathToFile := "data/d2"
	data, err := ioutil.ReadFile(pathToFile)
	powerSum := 0
	if err != nil {
		fmt.Printf("Read failed with: %s", err.Error())
		os.Exit(2)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		power, exists := GetGameSetPower(line)
		if !exists {
			fmt.Println("Failed to calculate power for this game")
			os.Exit(3)
		}
		powerSum += power
	}
	fmt.Printf("Total power sum is: %d", powerSum)
}
