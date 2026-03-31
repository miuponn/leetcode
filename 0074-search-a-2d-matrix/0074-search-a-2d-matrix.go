func searchMatrix(matrix [][]int, target int) bool {
    j := len(matrix[0]) - 1
    lo, hi := 0, len(matrix) - 1
    for lo <= hi {
        i := (lo + hi) / 2
        if target > matrix[i][j] {
            lo = i + 1
        } else if target < matrix[i][0] {
            hi = i - 1
        } else {
            break
        }
    }
    if lo > hi { return false }
    i := (lo + hi) / 2
    lo, hi = 0, j
    for lo <= hi {
        j := (lo + hi) / 2
        if target > matrix[i][j] {
            lo = j + 1
        } else if target < matrix[i][j] {
            hi = j - 1
        } else {
            return true
        }
    }
    return false
}