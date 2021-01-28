package main

var finishgame = false
var turncounts = 0
var gameboard = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
var key = "O"

//To be implemented
type Scoreboard struct {
	user     int
	computer int
}

func main() {
	StartGame()
}
