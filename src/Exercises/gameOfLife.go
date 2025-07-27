package main

type GameBoard []([]bool)

func PlayGameOfLife(initialBoard GameBoard, numGens int) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1])
	}

	return boards
}

func InitializeBoard(numRows, numCols int) GameBoard {
	board := make(GameBoard, numRows)

	for r := range board {
		board[r] = make([]bool, numCols)
	}

	return board
}

func UpdateBoard(currBoard GameBoard) GameBoard {
	numRows := CountRows(currBoard)
	numCols := CountCols(currBoard)
	board := InitializeBoard(numRows, numCols)

	for r := range currBoard {
		for c := range currBoard[r] {
			board[r][c] = UpdateCell(currBoard, r, c)
		}
	}

	return board
}

func UpdateCell(board GameBoard, r, c int) bool {
	numNeighbors := CountLiveNeighbors(board, r, c)

	if board[r][c] == true {
		if numNeighbors == 2 || numNeighbors == 3 {
			return true
		} else {
			return false
		}
	} else {
		if numNeighbors == 3 {
			return true
		} else {
			return false
		}
	}
}

func CountLiveNeighbors(board GameBoard, r, c int) int {
	count := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if (i != r || j != c) && InField(board, i, j) && board[i][j] {
				count++
			}
		}
	}
	return count
}

func InField(board GameBoard, i, j int) bool {
	if i < 0 || i >= CountRows(board) || j < 0 || j >= CountCols(board) {
		return false
	}
	return true
}

func CountRows(board GameBoard) int {
	return len(board)
}

func CountCols(board GameBoard) int {
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	return len(board[0])
}
