func numRescueBoats(people []int, limit int) int {
    sort.Ints(people)
    l:=0
    r:= len(people)-1
    c:=0
    for l <= r {
        if people[l] + people[r] <= limit {
            l++
        }
        r--
        c++
    }
    return c
}