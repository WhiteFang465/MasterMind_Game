package rps

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	

)

type Round struct {
	Winner int `json:"winner"`
	
}

func PlayRound(playerValue int) Round {

	fmt.Println(randomNoGenerator())
	var result Round
	result.Winner = playerValue
	
	return result
}
func randomNoGenerator()int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999-1000) + 1000
}
