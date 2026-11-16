func largestNumber(nums []int) string {
    arr := make([]string, len(nums))
    for i, num := range nums {
        arr[i] = strconv.Itoa(num)
    }

    sort.Slice(arr, func(i, j int) bool {
        return arr[i]+arr[j] > arr[j]+arr[i]
    })

    res := strings.Join(arr, "")
    if res[0] == '0' {
        return "0"
    }
    return res
}