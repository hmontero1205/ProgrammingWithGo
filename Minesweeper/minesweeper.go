package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
	"strings"
	"strconv"
)

var board [][] string
var boardView [][] bool
var currentSeed int64
var gameOver bool

func main() {
	fmt.Println("Let's play MineSweeper!!!\nWhat difficulty should we play on? (easy, medium, hard)")
	reader := bufio.NewReader(os.Stdin)
	diff,_ := reader.ReadString('\n')
	diff = strings.TrimSpace(diff)
	switch diff {
	case "easy":
		board = makeString2DArray(10, 10)
	case "medium":
		board = makeString2DArray(30, 30)
	case "hard":
		board = makeString2DArray(50, 50)
	default:
		board = makeString2DArray(30, 30)
	}
	boardView = makeBool2DArray(len(board), len(board[0]))

	constructBoard()
	playGame()
	//printBoard()
	
}

func playGame() {
	reader := bufio.NewReader(os.Stdin)
	for !checkWinner() && !gameOver {
		printBoard()
		fmt.Println("Please don't break me :)\nrow pls")
		inRow,_ := reader.ReadString('\n')
		fmt.Println("col pls")
		inCol,_ := reader.ReadString('\n')
		row,_ := strconv.Atoi(strings.TrimSpace(inRow))
		col,_ := strconv.Atoi(strings.TrimSpace(inCol))
		for boardView[row][col] {
			fmt.Println("You chose that already >:(\nrow pls")
			inRow,_ = reader.ReadString('\n')
			fmt.Println("col pls")
			inCol,_ = reader.ReadString('\n')
			row,_ = strconv.Atoi(strings.TrimSpace(inRow))
			col,_ = strconv.Atoi(strings.TrimSpace(inCol))
		}

		showTile(row, col)
		if board[row][col] == "X"{
			gameOver = true;
			printBoard();
			fmt.Println("You lost :'(")
		}


	}
}

func showTile(r, c int) {
	if !boardView[r][c]{
		boardView[r][c] = true;
		if board[r][c] == " " {
			if r - 1 >= 0 {
				showTile(r - 1, c)
			}
			if r + 1 < len(board) {
				showTile(r + 1, c)
			}
			if c + 1 < len(board[0]) {
				showTile(r, c + 1)
			}
			if c - 1 >= 0 {
				showTile(r, c - 1)
			}
		}
		return
	}
}

func makeString2DArray(rows int, cols int) [][]string {
	arr := make([][]string, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]string, cols)
	}

	return arr
}

func makeBool2DArray(rows int, cols int) [][]bool {
	arr := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]bool, cols)
	}

	return arr
}

func constructBoard() {
	for r :=0; r < len(board); r++ {
		for c :=0; c < len(board[r]); c++ {
			if random(1,21) < 3 {
				board[r][c] = "X"
			} else {
				board[r][c] = " "
			}
		}
	}

	for r :=0; r < len(board); r++ {
		for c :=0; c < len(board[r]); c++ {
			if board[r][c] != "X" {
				count := 0
				for sr := max(0, r - 1); sr < min(len(board), r + 2); sr++ {
					for sc:= max(0, c - 1); sc < min(len(board[sr]), c + 2); sc++ {
						if !(sr == r && sc == c) && board[sr][sc] == "X" {
							count++
						}
					}
				}
				if count > 0 {
					board[r][c] = strconv.Itoa(count)
				}
			}
		}
	}
}

func printBoard() {
	for row := 0; row < len(board) + 1; row++ {
		for col := 0; col < len(board[0]) + 1; col++ {
			if row == 0 {
				if col != len(board[0]) {
					fmt.Printf(strconv.Itoa(col) + " ")
					if col < 10 {
						fmt.Printf(" ")
					}
				}
			} else {
				if col == len(board[0]) {
					fmt.Printf(strconv.Itoa(row - 1) + "  ")
				} else {
					if boardView[row - 1][col] {
						fmt.Printf(board[row - 1][col] + "  ")
					} else {
						fmt.Printf("O  ")
					}

				}
			}

		}
		fmt.Println("")
	}
}

func random(min, max int) int {
    newSeed := int64(time.Now().UTC().UnixNano())
   	for newSeed == currentSeed {
   		newSeed = int64(time.Now().UnixNano())
   	}
   	rand.Seed(newSeed)
   	currentSeed = newSeed
    return rand.Intn(max - min) + min
}

func min(a, b int) int {
	if a < b {
		return a 
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a 
	}

	return b
}

func checkWinner() bool {
	for r := 0; r < len(board); r++ {
		for c:= 0; c < len(board[r]); c++ {
			if board[r][c] != "X" && !boardView[r][c] {
				return false
			}
		}
	}

	printBoard()
	fmt.Println("You won! :)")
	return true
}