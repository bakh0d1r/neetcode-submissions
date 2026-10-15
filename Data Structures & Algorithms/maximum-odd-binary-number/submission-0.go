func maximumOddBinaryNumber(s string) string {
    arr := []byte(s)
    left := 0

    for i := 0; i < len(arr); i++ {
        if arr[i] == '1' {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }

    arr[left-1], arr[len(arr)-1] = arr[len(arr)-1], arr[left-1]
    return string(arr)
}