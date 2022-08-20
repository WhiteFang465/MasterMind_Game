package rps

import (
	"fmt"
	"strconv"
	
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

func PlayRound(playerValue int,generatedNumber int) Round {
	var result Round
	
	if !digitValidator(playerValue) {
		
	if countDigits(playerValue)==4 {
		systemGenerated:=generatedNumber
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
		
	}else{
		result.Winner="Enter a 4 digit number"
		result.CountInPlace=strconv.Itoa(0)
		result.CountMissplace=strconv.Itoa(0)
		result.ComputerGenerated=0
		return result
	}
	
	}else{
	result.Winner="Every digit of the number should be unique"
	result.CountInPlace=strconv.Itoa(0)
	result.CountMissplace=strconv.Itoa(0)
	result.ComputerGenerated=0
	return result
}

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



func countDigits(number int) int {
	counter := 0
	for number != 0 {
		number = number / 10
		counter += 1
	}
	return counter
}

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
