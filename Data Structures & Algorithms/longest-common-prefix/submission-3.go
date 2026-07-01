func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
		return ""
	}
	arr:=[]byte(strs[0])
	for i:=1;i<len(strs);i++{
		if len(arr) == 0 || len(strs[i]) == 0{
			return ""
		}
		if len(arr) > len(strs[i]){
			arr=arr[:len(strs[i])]
		}
		for f,j:=0,0; f<len(arr) && j<len(strs[i]); {
			if arr[f] != strs[i][j] {
				arr=arr[:f]
				break
			}
			f++
			j++
		}
	}
	return string(arr)
}
