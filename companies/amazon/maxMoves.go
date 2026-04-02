package amazon

func maxMoves(turns int, board [8][8]string) int {
	king := [2]int{0, 0}
	for r, row := range board {
		for cl, piece := range row {
			if piece == "k" {
				king = [2]int{r, cl}
			}
		}
	}
	return countCaptures(king, board, turns)
}

var directions = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func countCaptures(k [2]int, b [8][8]string, t int) int {
	if t == 0 {
		return 0
	}

	best := 0
	for _, d := range directions {
		nr, nc := k[0]+d[0], k[1]+d[1]
		if nr < 0 || nr > 7 || nc < 0 || nc > 7 {
			continue
		}

		captured := 0
		newBoard := b // arrays are value types in Go — safe copy
		if newBoard[nr][nc] == "x" {
			captured = 1
			newBoard[nr][nc] = "" // remove captured piece
		}

		result := captured + countCaptures([2]int{nr, nc}, newBoard, t-1)
		if result > best {
			best = result
		}
	}
	return best
}
