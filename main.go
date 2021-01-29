package main

import (
	"log"
	"os")


var finishgame = false
var turncounts = 0
var gameboard = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
var key = "O"

func init() {
	userScore := os.Getenv("userScore")
	computerScore := os.Getenv("computerScore")
	if userScore==""&&computerScore==""{
	err1 := os.Setenv("userScore","0")
	err2 := os.Setenv("computerScore","0")
	
	if err1!=nil||err2!=nil{
		log.Fatal("Could not set up the game,Please Restart")
	}
    }
}

func main() {
	StartGame()
}
