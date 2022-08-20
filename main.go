package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
    "myapp/Backend"
	"net/http"
	"strconv"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "./FrontEnd/index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	
	result := rps.PlayRound(playerChoice, generatedNumber)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}


var generatedNumber int

func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	generatedNumber = uniqueGenerator()
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
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
func uniqueGenerator() int {

	var uniqueNumber int

	for digitValidator(uniqueNumber) {
		uniqueNumber = randomNoGenerator() //1234

	}

	return uniqueNumber
}
func randomNoGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999-1000) + 1000
}
