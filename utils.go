package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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
	case "" :
		fmt.Println("Please re-enter your option")
		time.Sleep(time.Second * 2)
		fmt.Println("\n")
		startGame()
	default:
		fmt.Println("Please re-enter your option")
		time.Sleep(time.Second * 2)
		fmt.Println("\n")
		startGame()
	}
}
