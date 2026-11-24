func maxVowels(s string, k int) int {
    vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

    l, cnt, res := 0, 0, 0
    for r := 0; r < len(s); r++ {
        if vowels[s[r]] {
            cnt++
        }
        if r-l+1 > k {
            if vowels[s[l]] {
                cnt--
            }
            l++
        }
        res = max(res, cnt)
    }
    return res
}