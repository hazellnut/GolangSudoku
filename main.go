package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println(quote.Go())

	// var doggy Dog
	// doggy.Animal = Animal{"Jason", 4}
	// doggy.floof_score = 10

	// fmt.Printf("This is a go test string, the dog's name is %s and he is so floofy: %d", doggy.Animal.name, doggy.floof_score)

	// router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.Run("localhost:8080")
	example_sudoku := [9][9]int{
		{1, 0, 0, 4, 8, 9, 0, 0, 6},
		{7, 3, 0, 0, 0, 0, 0, 4, 0},
		{0, 0, 0, 0, 0, 1, 2, 9, 5},
		{0, 0, 7, 1, 2, 0, 6, 0, 0},
		{5, 0, 0, 7, 0, 3, 0, 0, 8},
		{0, 0, 6, 0, 9, 5, 7, 0, 0},
		{9, 1, 4, 6, 0, 0, 0, 0, 0},
		{0, 2, 0, 0, 0, 0, 0, 3, 7},
		{8, 0, 0, 5, 1, 2, 0, 0, 4}}

	new_arr := solveSudoku(example_sudoku)
	print_sudoku(new_arr)

}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type Animal struct {
	name string
	size int
}

type Dog struct {
	Animal
	floof_score int
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postSudoku(c *gin.Context) {
	solveSudoku(arr)
}

func print_sudoku(arr [9][9]int) {
	str := ""
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			str += fmt.Sprintf("%d", arr[i][j])
		}
		str += "\n"
	}
	fmt.Printf(str)
}

func solveSudoku(arr [9][9]int) [9][9]int {
	//logic for solving sudoku:
	//take the first input cell that is empty (i.e. 0) and test from 1 to 9
	//if the value is valid, test the next empty cell
	//if there is no valid value for all combinations, return false
	//
	r0, c0 := findNextEmptyCell(arr, 0, 0)
	for i := 1; i <= 9; i++ {
		arr[r0][c0] = i
		if solveCell(&arr, 0, 0) {
			return arr
		}
	}

	return arr
}

func solveCell(arr *[9][9]int, row int, col int) bool {
	//immediately return false if there is a clash with rows and columns
	if !checkRowsAndColumns(arr, row, col) || (arr[row][col] == 0) {
		return false
	}

	new_row, new_col := findNextEmptyCell(*arr, row, col)

	//exit condition - we have reached the end of rows and columns that are not filled
	if new_row == -1 && new_col == -1 {
		return true
	}
	for i := 1; i <= 9; i++ {
		arr[new_row][new_col] = i
		if solveCell(arr, new_row, new_col) {
			return true
		}
	}
	//if we reach this statement we did not solve for the current cell, reset the next level cell (to 0) and return false
	arr[new_row][new_col] = 0
	return false

}

func checkRowsAndColumns(arr *[9][9]int, row int, col int) bool {
	value := arr[row][col]
	for i := 0; i < 9; i++ {
		if i == row {
			continue
		}
		if arr[i][col] == value {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if i == col {
			continue
		}
		if arr[row][i] == value {
			return false
		}
	}
	return true

}

func findNextEmptyCell(arr [9][9]int, row int, col int) (int, int) {
	for i := col; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if i == col && j < row {
				continue
			}
			if arr[j][i] == 0 {
				return j, i
			}
		}
	}
	return -1, -1
}
