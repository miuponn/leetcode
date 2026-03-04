func isValidSudoku(board [][]byte) bool {
    rows := make(map[int]map[byte]bool)
    cols := make(map[int]map[byte]bool)
    sqrs := make(map[int]map[byte]bool)
    
    for i, row := range board{
        if rows[i] == nil { rows[i] = make(map[byte]bool) }
        for j, digit := range row {
            if cols[j] == nil { cols[j] = make(map[byte]bool) }
            k := (i/3) * 3 + (j/3)
            if sqrs[k] == nil { sqrs[k] = make(map[byte]bool) }
            if digit == '.' { continue }
            if rows[i][digit] || cols[j][digit] || sqrs[k][digit] {
                return false
            } else {
                rows[i][digit] = true
                cols[j][digit] = true
                sqrs[k][digit] = true
            }
        }
    }
    return true
}
