package advent

import (
	"bufio"
	"fmt"
	"os"
)

func DayThree(file string) int {
	readfile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	prioritySum := 0

	fileScanner := bufio.NewScanner(readfile)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		itemMap := make(map[rune]bool)

		for _, item := range line[0 : len(line)/2] {
			itemMap[item] = true
		}

		for _, item := range line[len(line)/2:] {
			if itemMap[item] {
				prioritySum += getItemPriority(item)
				break
			}
		}
	}

	return prioritySum
}

func getItemPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 96
	}
	return int(item) - 38
}
