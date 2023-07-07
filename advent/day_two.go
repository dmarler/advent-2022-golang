package advent

import (
	"bufio"
	"fmt"
	"os"
)

var DayTwoPartTwo = true

func DayTwo(file string) int {

	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	score := 0

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		score += getLineValue(fileScanner.Text())
	}

	return score
}

func getLineValue(line string) int {
	// Note: I expect there's more clever ways to calculate the numbers,
	//		 but with the fact that there's only 18 options to map out
	//		 and the fact that I needed to learn the syntax of switch cases in go,
	//		 I went with this option.
	if DayTwoPartTwo {
		switch line {
		case "A X":
			return 3
		case "A Y":
			return 4
		case "A Z":
			return 8
		case "B X":
			return 1
		case "B Y":
			return 5
		case "B Z":
			return 9
		case "C X":
			return 2
		case "C Y":
			return 6
		case "C Z":
			return 7
		}
	} else {
		switch line {
		case "A X":
			return 4
		case "A Y":
			return 8
		case "A Z":
			return 3
		case "B X":
			return 1
		case "B Y":
			return 5
		case "B Z":
			return 9
		case "C X":
			return 7
		case "C Y":
			return 2
		case "C Z":
			return 6
		}
	}
	return 0
}
