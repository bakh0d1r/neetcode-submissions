func mergeAlternately(word1 string, word2 string) string {
	str:=[]byte{}
	l,r := 0,0
	for l < len(word1) && r < len(word2) {
		str=append(str,word1[l],word2[r])
		l++
		r++
	}
	str=append(str,word1[l:]...)
	str=append(str,word2[r:]...)
	return string(str)
}

