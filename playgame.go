package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Printboard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("|", gameboard[i][j])
		}
		fmt.Print("|\n")
	}
	fmt.Println()
}

func StartGame() {
	var turn string
	fmt.Scanf("%s", &key)
	if key == "X" {
		UserPlay()
		turn = "computer"
	} else if key == "O" {
		//if user chooses O then computer will start
		ComputerPlay()
		turn = "player"
	} else {
		fmt.Println("Invalid Input")
		return
	}

	for turncounts < 9 {
		Printboard()
		var done bool
		if turn == "player" {
			done = UserPlay()
			if !done {
				break
			}
			ok := checkForWin()
			if ok {
				Printboard()
				UpdateScore(turn)
				fmt.Println("User Wins")
				return
			}
		} else {
			ComputerPlay()
			ok := checkForWin()
			if ok {
				UpdateScore(turn)
				Printboard()
				fmt.Println("Computer Wins")
				return
			}
		}
		switchTurn(&turn)
		// Printboard()
	}
	if turncounts == 9 {
		fmt.Println("The Game is a Draw")
	}
}

func UserPlay() bool {
	// fmt.Println("user playing")
	var done = false
	for !done {
		var move, points string
		fmt.Scan(&move)
		if move == "quit" {
			fmt.Println("Computer Wins")
			turncounts = 10
			return done
		}
		if move == "scoreboard"{
			PrintScoreBoard()
			fmt.Scan(&move)
		}
		
		fmt.Scanf("%s", &points)
		pos1, _ := strconv.Atoi(string(points[0]))
		pos2, _ := strconv.Atoi(string(points[2]))
		if gameboard[pos1][pos2] == " " {
			gameboard[pos1][pos2] = key
			done = true
		} else {
			fmt.Println("Invalid move, place already taken")
		}
	}
	turncounts++
	switchKey(&key)
	return done
}

func ComputerPlay() {
	// fmt.Println("computer playing")
	var done = false
	rand.Seed(time.Now().UnixNano())
	for !done {
		
		pos1 := rand.Intn(2)
		pos2 := rand.Intn(2)
		if gameboard[pos1][pos2] == " " {
			gameboard[pos1][pos2] = key
			done = true
		}
	}
	turncounts++
	switchKey(&key)
}

func switchKey(k *string) {
	if *k == "X" {
		*k = "O"
	} else {
		*k = "X"
	}
}

func switchTurn(t *string) {
	if *t == "player" {
		*t = "computer"
	} else {
		*t = "player"
	}
}

func checkForWin() bool {
	//compare key and ascii value
	board := ParseGameBoard()
	for i := range board {
		if board[i] == "XXX" || board[i] == "OOO" {
			return true
		}
	}
	return false
}

func ParseGameBoard() [8]string {
	var b [8]string
	b[0] = gameboard[0][0] + gameboard[0][1] + gameboard[0][2]
	b[1] = gameboard[1][0] + gameboard[1][1] + gameboard[1][2]
	b[2] = gameboard[2][0] + gameboard[2][1] + gameboard[2][2]
	b[3] = gameboard[0][0] + gameboard[1][0] + gameboard[2][0]
	b[4] = gameboard[0][1] + gameboard[1][1] + gameboard[2][1]
	b[5] = gameboard[0][2] + gameboard[1][2] + gameboard[2][2]
	b[6] = gameboard[0][0] + gameboard[1][1] + gameboard[2][2]
	b[7] = gameboard[0][2] + gameboard[1][1] + gameboard[2][0]
	return b
}

func UpdateScore(turn string) {
	userScore := os.Getenv("userScore")
	computerScore := os.Getenv("computerScore")
	fmt.Println(userScore,computerScore)
	if turn == "computer" {
		  Score,_ :=strconv.Atoi(computerScore)
		  Score++
		  err:=os.Setenv("computerScore",strconv.Itoa(Score))
		  if err!=nil{
			  fmt.Println("error while updating score")
		  }

	}else{
		Score,_ :=strconv.Atoi(userScore)
		  Score++
		  err:=os.Setenv("userScore",strconv.Itoa(Score))
		  if err!=nil{
			  fmt.Println("error while updating score")
		  }
	}

}

func PrintScoreBoard(){
	userScore := os.Getenv("userScore")
	computerScore := os.Getenv("computerScore")
	fmt.Println("| Computer | User |")
	fmt.Println("|      ",computerScore," |"," ",userScore," |")
}