package trees

func numTrees(n int) int {
    if n < 2 {
        return 1
    }
    if n == 2 {
        return 2
    }
    count := 0
    for i:=1; i <= n; i++ {
        count += numTrees(i-1) * numTrees(n-i)
    }
    return count
}