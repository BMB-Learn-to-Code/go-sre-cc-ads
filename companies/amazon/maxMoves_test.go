package amazon

import "testing"

func TestMaxMoves(t *testing.T) {
	var board [8][8]string

	// King at [4][4], pieces at [3][4] (adjacent above), [3][2], [6][4]
	//
	// . . . . . . . .   row 0
	// . . . . . . . .   row 1
	// . . . . . . . .   row 2
	// . . x . x . . .   row 3: x at col 2, x at col 4
	// . . . . k . . .   row 4: k at col 4
	// . . . . . . . .   row 5
	// . . . . x . . .   row 6: x at col 4
	// . . . . . . . .   row 7

	board[4][4] = "k"
	board[3][4] = "x"
	board[6][4] = "x"
	board[3][2] = "x"

	// Best path: [4][4] -> [3][4](cap) -> [3][3] -> [3][2](cap) = 2 captures
	got := maxMoves(3, board)
	want := 2
	if got != want {
		t.Errorf("maxMoves(3, board) = %d, want %d", got, want)
	}
}
