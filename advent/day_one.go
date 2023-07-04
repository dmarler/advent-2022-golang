package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DayOne(file string) int {

	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []int
	var maxCaloriesArr [3]int
	currentCalories := 0

	for fileScanner.Scan() {
		line, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			line = 0
		}
		fileLines = append(fileLines, line)
	}

	for _, calories := range fileLines {
		if calories == 0 {

			if currentCalories > maxCaloriesArr[0] {
				maxCaloriesArr[2] = maxCaloriesArr[1]
				maxCaloriesArr[1] = maxCaloriesArr[0]
				maxCaloriesArr[0] = currentCalories
			} else if currentCalories > maxCaloriesArr[1] {
				maxCaloriesArr[2] = maxCaloriesArr[1]
				maxCaloriesArr[1] = currentCalories
			} else if currentCalories > maxCaloriesArr[2] {
				maxCaloriesArr[2] = currentCalories
			}

			currentCalories = 0

			continue
		}
		currentCalories += calories
	}

	fmt.Printf("Part Two: %d\n", maxCaloriesArr[0]+maxCaloriesArr[1]+maxCaloriesArr[2])

	return maxCaloriesArr[0]
}
