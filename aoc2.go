package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var maxColoredCubes = map[string]int{"red": 12, "green": 13, "blue": 14}

func GetPossibleGame(line string) (bool, int) {
	gameNumber := 0
	possibleGame := false
	game := strings.Split(line, ":")
	number, found := strings.CutPrefix(game[0], "Game ")
	if !found {
		return false, 0
	}
	gameNumber, _ = strconv.Atoi(number)
	sets := strings.Split(game[1], ";")
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			cubeData := strings.Split(cube, " ")
			drawCount, _ := strconv.Atoi(cubeData[0])
			if drawCount > maxColoredCubes[cubeData[1]] {
				return false, 0
			}
		}
	}
	possibleGame = true
	return possibleGame, gameNumber
}
func main() {

	pathToFile := "data/d2"
	data, err := ioutil.ReadFile(pathToFile)
	gameSum := 0
	if err != nil {
		fmt.Printf("Read failed with: %s", err.Error())
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		possibleGame, gameNumber := GetPossibleGame(line)
		if possibleGame && gameNumber > 0 {
			gameSum += gameNumber
		}
	}
	fmt.Printf("Total game sum is: %d", gameSum)
}

/*
	Split data line by line
	Split data based on colon delimiter to tell the game number
	Split data based on semicolon delimiter to set type
	Split data based on comma to tell cube color
*/
