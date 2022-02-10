package rps

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	 ARRAY_SIZE=100 
	 

)

type Round struct {
	Winner int `json:"winner"`
	
}

func PlayRound(playerValue int) Round {

	fmt.Println(uniqueGenerator())
	var result Round
	result.Winner = playerValue
	
	return result
}
func randomNoGenerator()int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999-1000) + 1000
}

func digitValidator (number int) bool{
	 firstNumber  := number / 1000;
     secondNumber := (number / 100) - (firstNumber * 10);
     thirdNumber  := (number % 100) / 10;
     fourthNumber := number % 10;

	 if ((firstNumber == secondNumber) || (firstNumber == thirdNumber) || (firstNumber == fourthNumber)) {
		 return true;
	 }else if ((secondNumber == thirdNumber) || (secondNumber == fourthNumber)) {
		 return true;
	 }else if (thirdNumber == fourthNumber) {
		 return true;
	 }
 
	 return false;

}

func uniqueGenerator()int{
	
	var uniqueNumber int

	for flag:=true;flag;flag=false  {
		for i := 0; i < ARRAY_SIZE; i++ {
			for digitValidator(uniqueNumber){
					uniqueNumber=randomNoGenerator()
			}
			flag=true
			
		}
	}
	return uniqueNumber
}

func countDigits(number int)int{
	counter:=0
	for(number!=0){
number=number/10
counter+=1
	}
	return counter
}

