package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools

type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"

type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	chosen, ok := cb[rank]
	if !ok{
		return 0
	}
	i := 0
	for _, j := range chosen{
		if j{
			i++
		}
	}
	return i
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}
	fileCount := 0
	for _, rank := range cb {
		if rank[file-1] {
			fileCount += 1
		}
	}
	return fileCount
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	squares := 0
	for _, rank := range cb {
		for range rank {
			squares += 1
		}
	}
	return squares
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	occupiedSquares := 0
	for _, rank := range cb {
		for _, isOnBoard := range rank {
			if isOnBoard {
				occupiedSquares += 1
			}
		}
	}
	return occupiedSquares
}
