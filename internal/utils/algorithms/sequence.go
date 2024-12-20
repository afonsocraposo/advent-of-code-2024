package algorithms

// dpSequenceCheck uses dynamic programming to check if the passed `sequence` can be constructed with the `pieces` passed.
func DpSequenceCheck(sequence string, pieces []string) bool {
	n := len(sequence)
	dp := make([]bool, n+1)
	dp[0] = true // base case: empty sequence can be constructed

	for i := 1; i < n+1; i++ {
		dp[i] = false
		for _, piece := range pieces {
			piece_len := len(piece)
			if i >= piece_len && sequence[i-piece_len:i] == piece {
				dp[i] = dp[i] || dp[i-piece_len]
			}
		}
	}
	return dp[n]
}

func DpSequenceArrangementsCount(sequence string, pieces []string) int {
	n := len(sequence)
	dp := make([]int, n+1)
	dp[0] = 1 // base case: one way to construct an empty sequence

	for i := 1; i < n+1; i++ {
		for _, piece := range pieces {
			piece_len := len(piece)
			if i >= len(piece) && sequence[i-piece_len:i] == piece {
				dp[i] += dp[i-piece_len]
			}
		}
	}
	return dp[n]
}

type stateRef struct {
	PrevState int
	Piece     string
}

func dpSequenceArrangements(sequence string, pieces []string) [][]stateRef {
	n := len(sequence)
	dp := make([][]stateRef, n+1)
	dp[0] = []stateRef{}

	for i := 1; i < n+1; i++ {
		dp[i] = []stateRef{}
		for _, piece := range pieces {
			piece_len := len(piece)
			if i >= piece_len && sequence[i-piece_len:i] == piece {
				// Add a reference to the previous state and the current piece
				dp[i] = append(dp[i], stateRef{i - piece_len, piece})
			}
		}
	}
	return dp
}

func reconstruct(dp [][]stateRef, index int) [][]string {
	if index == 0 {
		return [][]string{{}}
	}

	arrangements := [][]string{}
	for _, s := range dp[index] {
		prevInvex := s.PrevState
		piece := s.Piece
		for _, arrangement := range reconstruct(dp, prevInvex) {
			a := append(append([]string{}, arrangement...), piece)
			arrangements = append(arrangements, a)
		}
	}
	return arrangements
}

func DpSequenceArrangements(sequence string, pieces []string) [][]string {
	n := len(sequence)
	dp := dpSequenceArrangements(sequence, pieces)
	return reconstruct(dp, n)
}
