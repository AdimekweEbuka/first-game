package main

import (
	"fmt"
	"math/rand"
	"time"
)

const totalQuestions = 5

//GENERATES COMMENT AFTER AN ANSWER IS GOTTEN CORRECTLY
func RandomWTags() string {
	WinnerTags1 := "BOSSMAN, FIRE DOWN"
	WinnerTags2 := "CHAIRMO, AFTER YOU NAH YOU OH"
	WinnerTags3 := "ALAYE, CUT SOAP FOR ME NAH"
	WinnerTags4 := "OVERALL BEST IN SENSE!!!"

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	number := r1.Intn(500)
	err := "Please revisit WTag function"

	if number >= 0 && number <= 200 {
		return WinnerTags1
	} else if number >= 200 && number <= 300 {
		return WinnerTags2
	} else if number >= 300 && number <= 400 {
		return WinnerTags3
	} else if number >= 400 && number <= 500 {
		return WinnerTags4
	}
	return err

}

//GENERATES COMMENT AFTER AN ANSWER IS GOTTEN INCORRECTLY
func RandomLTags() string {
	LoserTags1 := "SEE AS YOU RESEMBLE CHELSEA FAN, GAME OVER"
	LoserTags2 := "IF YOU NO LEAVE THIS GAME NOW EHN, I GO COMOT MY SHOE CHASE YOU"
	LoserTags3 := "WEREY DEY DISGUISE LMAO, COMOT MY GAME ABEG"
	LoserTags4 := "ALAYE YAKUBU SEF NO DEY MISS REACH YOU"

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	number := r1.Intn(500)
	err := "Please revisit LTag function"

	if number >= 0 && number <= 200 {
		return LoserTags1
	} else if number >= 200 && number <= 300 {
		return LoserTags2
	} else if number >= 300 && number <= 400 {
		return LoserTags3
	} else if number >= 400 && number <= 500 {
		return LoserTags4
	}
	return err

}

func Lawrence() {
	fmt.Println("You chose Lawrence!")

	question1 := "How old is lawrence?"
	answer1 := "35"
	ScoreOne := Scanner(question1, answer1)

	question2 := "Lawrence get babe?"
	answer2 := "NO"
	ScoreTwo := Scanner(question2, answer2)

	finalScore := ScoreOne + ScoreTwo
	fmt.Println("Your final score is ", finalScore, "/", totalQuestions)
}

func Kennedy() {
	fmt.Println("You chose Kennedy!")

	question1 := "Who's Kennedys best friend"
	answer1 := "FEGOR"
	ScoreOne := Scanner(question1, answer1)

	question2 := "What is Kennedys body count"
	answer2 := "40"
	ScoreTwo := Scanner(question2, answer2)

	finalScore := ScoreOne + ScoreTwo
	fmt.Println("Your final score is ", finalScore, "/", totalQuestions)

}
