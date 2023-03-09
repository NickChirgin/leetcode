package binarysearch

func maximumRemovals(s string, p string, removable []int) int {
    count := 0
    removed := map[int]struct{}{}
    var check func(s,p string) bool
    check = func (s, p string) bool {
        i1, i2 := 0, 0
        for i1 < len(s) && i2 < len(p) {
            if _, ok := removed[i1]; ok || s[i1] != p[i2] {
                i1++
                continue
            }
            i1++
            i2++
        }
        return i2 == len(p)
    }
    l,r := 0, len(removable) -1
    for l <= r {
        m := (l+r) / 2
        removed = map[int]struct{}{}
        for _, v := range removable[:m+1] {
            removed[v] = struct{}{}
        }
        if check(s, p) {
            count = max(count, m+1)
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return count
}

func max(a, b int) int{
    if a > b {
        return a
    }
    return b
}