package main

import (
	"fmt"
)
const totalQuestions = 5

func Lawrence() {
	fmt.Println("You chose Lawrence!")

	question1 := "How old is lawrence?"
	answer1 := "35"
	ScoreOne:= Scanner(question1, answer1)
	
	
	question2 := "Lawrence get babe?"
	answer2 := "NO"
	ScoreTwo:= Scanner(question2, answer2)
	
	finalScore:= ScoreOne + ScoreTwo 
	fmt.Println("Your final score is ", finalScore,"/", totalQuestions )
}

func Kennedy() {
	fmt.Println("You chose Kennedy!")

	question1 := "Who's Kennedys best friend"
	answer1 := "FEGOR"
	ScoreOne:= Scanner(question1, answer1)
	

	question2 := "What is Kennedys body count"
	answer2 := "40"
	ScoreTwo:= Scanner(question2, answer2)

	finalScore:= ScoreOne + ScoreTwo 
	fmt.Println("Your final score is ", finalScore,"/", totalQuestions )
	
}
