package main

import (
	"strconv"
)

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// GameBoard is a two-dimensional slice of ints
type GameBoardAutomaton [][]int

func PlayAutomaton(numGens int, neighborhoodType string, initialBoard GameBoardAutomaton, ruleStrings map[string]int) []GameBoardAutomaton {
	boards := make([]GameBoardAutomaton, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoardAutomaton(boards[i-1], neighborhoodType, ruleStrings)
	}

	return boards
}

func UpdateBoardAutomaton(currBoard GameBoardAutomaton, neighborhoodType string, ruleStrings map[string]int) GameBoardAutomaton {
	numRows := CountRowsAutomaton(currBoard)
	numCols := CountColsAutomaton(currBoard)
	newBoard := InitializeBoardAutomaton(numRows, numCols)

	for r := range currBoard {
		for c := 0; c < numCols; c++ {
			newBoard[r][c] = UpdateCellAutomaton(currBoard, r, c, neighborhoodType, ruleStrings)
		}
	}
	return newBoard
}

// InitializeBoard takes a number of rows and columns as inputs and returns a gameboard with appropriate number of rows and colums, where all values = 0.
func InitializeBoardAutomaton(numRows, numCols int) GameBoardAutomaton {
	// make a 2-D slice (default values = false)
	var board GameBoardAutomaton
	board = make(GameBoardAutomaton, numRows)
	// now we need to make the rows too
	for r := range board {
		board[r] = make([]int, numCols)
	}

	return board
}

func UpdateCellAutomaton(board GameBoardAutomaton, row, col int, neighborhoodType string, ruleStrings map[string]int) int {
	neighborhood := NeighborhoodToString(board, row, col, neighborhoodType)
	rule, exists := ruleStrings[neighborhood]

	if exists {
		return rule
	}

	return 0
}

func NeighborhoodToString(board GameBoardAutomaton, row, col int, neighborhoodType string) string {
	neighborhood := strconv.Itoa(board[row][col])
	var offsets [][2]int

	if neighborhoodType == "Moore" {
		offsets = [][2]int{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, 1}, {1, 1}, {1, 0},
			{1, -1}, {0, -1},
		}
	} else if neighborhoodType == "vonNeumann" {
		offsets = [][2]int{
			{-1, 0}, {0, 1}, {1, 0}, {0, -1},
		}
	}

	for _, offset := range offsets {
		x := offset[0]
		y := offset[1]
		if InFieldAutomaton(board, row+x, col+y) {
			neighborhood += strconv.Itoa(board[row+x][col+y])
		} else {
			neighborhood += "0"
		}
	}

	return neighborhood
}

func InFieldAutomaton(board GameBoardAutomaton, i, j int) bool {
	if i < 0 || i >= CountRowsAutomaton(board) || j < 0 || j >= CountColsAutomaton(board) {
		return false
	}
	return true
}

func CountRowsAutomaton(board GameBoardAutomaton) int {
	return len(board)
}

func CountColsAutomaton(board GameBoardAutomaton) int {
	if CountRowsAutomaton(board) == 0 {
		panic("Error: empty board given to CountCols")
	}

	return len(board[0])
}
