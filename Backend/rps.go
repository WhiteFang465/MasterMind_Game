package rps

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	ARRAY_SIZE = 100
)

type Round struct {
	Winner         string `json:"winner"`
	CountMissplace string `json:"miss_place"`
	CountInPlace   string `json:"in_place"`
	ComputerGenerated int `json:"computer_generated"`
}

func PlayRound(playerValue int) Round {
	var result Round
	
	systemGenerated := uniqueGenerator()
	counterInPlacement := inPlaceMatches(playerValue, systemGenerated)
	counterOutOfPlacement := outOfPlace(playerValue, systemGenerated)
	fmt.Println(systemGenerated)
	if counterInPlacement == 4 {
		result.Winner = "You Win"
		result.CountInPlace = strconv.Itoa(counterInPlacement)
		result.CountMissplace = strconv.Itoa(counterOutOfPlacement)
		result.ComputerGenerated=systemGenerated
		return result
	} else {
		result.Winner = "Incorrect Number"
		result.CountInPlace = strconv.Itoa(counterInPlacement)
		result.CountMissplace = strconv.Itoa(counterOutOfPlacement)
		result.ComputerGenerated=systemGenerated
		return result
	}

}

func randomNoGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999-1000) + 1000
}

func digitValidator(number int) bool {
	firstNumber := number / 1000
	secondNumber := (number / 100) - (firstNumber * 10)
	thirdNumber := (number % 100) / 10
	fourthNumber := number % 10

	if (firstNumber == secondNumber) || (firstNumber == thirdNumber) || (firstNumber == fourthNumber) {
		return true
	} else if (secondNumber == thirdNumber) || (secondNumber == fourthNumber) {
		return true
	} else if thirdNumber == fourthNumber {
		return true
	}

	return false

}

func uniqueGenerator() int {

	var uniqueNumber int

	for flag := true; flag; flag = false {
		for i := 0; i < ARRAY_SIZE; i++ {
			for digitValidator(uniqueNumber) {
				uniqueNumber = randomNoGenerator()
			}
			flag = true

		}
	}
	return uniqueNumber
}

// func countDigits(number int) int {
// 	counter := 0
// 	for number != 0 {
// 		number = number / 10
// 		counter += 1
// 	}
// 	return counter
// }

func userInputNumberSeparator(number int) [4]int {
	var arr [4]int

	for i := 3; i >= 0; i-- {
		temp := number % 10
		number /= 10
		arr[i] = temp

	}
	return arr
}
func outOfPlace(userNumber int, generatedNumber int) int {

	var userInput [4]int
	var generatedNumbers [4]int
	userInput = userInputNumberSeparator(userNumber)
	generatedNumbers = userInputNumberSeparator(generatedNumber)
	count := 0
	for i := 0; i < 4; i++ {
		for j := 3; j >= 0; j-- {
			if i != j {
				if userInput[i] == generatedNumbers[j] {
					count++
					break
				}

			}
		}
	}
	return count

}

func inPlaceMatches(userNumber int, generatedNumber int) int {

	count := 0
	var userInput [4]int
	var generatedNumbers [4]int
	userInput = userInputNumberSeparator(userNumber)
	generatedNumbers = userInputNumberSeparator(generatedNumber)
	for i := 0; i < 4; i++ {
		if userInput[i] == generatedNumbers[i] {
			count++
		}

	}
	return count
}
