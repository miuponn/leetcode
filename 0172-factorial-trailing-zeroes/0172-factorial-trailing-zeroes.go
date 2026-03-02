func trailingZeroes(n int) int {
    k := 0                                  // k = num of trailing zeroes
    for i := 1; i <= n; i++ {               // iterative log 5 n 
        k += n/5                            // k = floor(n^i) until floor(n/5) = 5
        n = n/5                             
    }
    return k
}

// O(log n) iterative factorial solution, can be O(n) if hardcode if statements until n > 10,000, but O(log n) works for n > 10^4