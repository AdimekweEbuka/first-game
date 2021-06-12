package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

//SCANS USER INPUT
func Scanner(q1, a1 string) int {

	time.Sleep(time.Second * 2)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(q1)
	scanner.Scan()
	answer := scanner.Text()
	answer = strings.ToUpper(answer)

	if answer == a1 {
		fmt.Println(RandomWTags())
	} else {
		fmt.Println(RandomLTags())
	}
	//SCORING SECTION
	Score := 0
	if answer == a1 {
		Score++
	}
	return Score
}

//STARTS THE GAME
func startGame() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(`WHICH OF THE BARRISTERS DO YOU KNOW BEST?
	A) KENNEDY
	B) LAWRENCE`)
	scanner.Scan()
	input := scanner.Text()
	input = strings.ToUpper(input)

	switch input {
	case "A":
		Kennedy()
	case "B":
		Lawrence()
	default:
		fmt.Println("Please enter your option")
	}
}
