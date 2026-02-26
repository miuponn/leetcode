func romanToInt(s string) int {
    roman_num := map[byte]int{
    'I' : 1,
    'V' : 5,
    'X' : 10,
    'L' : 50,
    'C' : 100,
    'D' : 500,
    'M' : 1000,
    }
    res := roman_num[s[len(s)-1]]

    for i := (len(s) - 2); i >= 0; i-- {
        if roman_num[s[i]] < roman_num[s[i + 1]] {
            res -= roman_num[s[i]]
        } else {
            res += roman_num[s[i]]
        }
    }
    return res
}