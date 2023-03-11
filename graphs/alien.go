package graphs

func isAlienSorted(words []string, order string) bool {
    charMap := map[byte]int{}
    for i := range order {
        charMap[order[i]] = i
    }
    for j:=0; j < len(words)- 1; j++ {
        for i := range words[j] {
            if i == len(words[j+1]) {
                return false
            }
            if words[j][i] != words[j+1][i] {
                if charMap[words[j][i]] > charMap[words[j+1][i]] {
                    return false
                }
                break
            }
        }
    }
    return true
}