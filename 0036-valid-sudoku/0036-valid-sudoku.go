func isValidSudoku(board [][]byte) bool {
    // nested hashmaps with index:digit:bool
    rows := make(map[int]map[byte]bool)
    cols := make(map[int]map[byte]bool)
    sqrs := make(map[int]map[byte]bool)

    for i, row := range board{                                        // iterate over rows
        if rows[i] == nil { rows[i] = make(map[byte]bool) }           // init nested map for column
        for j, digit := range row {                                   // iterate over columns in each row
            if cols[j] == nil { cols[j] = make(map[byte]bool) }       // init nested map for column
            k := (i/3) * 3 + (j/3)                                    // calculate square index based on row, column
            if sqrs[k] == nil { sqrs[k] = make(map[byte]bool) }       // init nested map for square
            if digit == '.' { continue }                              // skip '.' entries
            if rows[i][digit] || cols[j][digit] || sqrs[k][digit] {   // check if digit in rows, cols, squares
                return false
            } else {                                                  // add digit to index:digit for row, col, square
                rows[i][digit] = true
                cols[j][digit] = true
                sqrs[k][digit] = true
            }
        }
    }
    return true
}
